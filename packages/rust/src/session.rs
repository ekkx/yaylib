use std::collections::HashMap;
use std::path::{Path, PathBuf};

use async_trait::async_trait;
use chrono::{DateTime, Utc};
use serde::{Deserialize, Serialize};
use tokio::sync::Mutex;

/// A persisted login session (tokens, device UUID, and metadata).
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Session {
    #[serde(default, skip_serializing_if = "String::is_empty")]
    pub email: String,
    pub user_id: i64,
    pub access_token: String,
    pub refresh_token: String,
    #[serde(default, skip_serializing_if = "String::is_empty")]
    pub device_uuid: String,
    /// API key used to refresh this session's tokens. Empty means the
    /// default app key; set for sessions obtained via the browser-login
    /// fallback (§6.2). Absent on older session files, which restore as
    /// app sessions.
    #[serde(default, skip_serializing_if = "String::is_empty")]
    pub oauth_key: String,
    pub updated_at: DateTime<Utc>,
}

/// Immutable snapshot of access and refresh tokens.
#[derive(Debug, Clone)]
pub struct Tokens {
    pub access: String,
    pub refresh: String,
}

#[derive(Debug, thiserror::Error)]
pub enum SessionError {
    #[error("no session for that email")]
    NoSession,

    #[error("email is required")]
    EmailRequired,

    #[error("session store not configured")]
    NotConfigured,

    #[error(transparent)]
    Io(#[from] std::io::Error),

    #[error("decode {path}: {source}")]
    Decode {
        path: String,
        #[source]
        source: serde_json::Error,
    },

    #[error(transparent)]
    Serialize(#[from] serde_json::Error),
}

/// Async trait for persisting login sessions.
#[async_trait]
pub trait SessionStore: Send + Sync {
    async fn load(&self, email: &str) -> Result<Session, SessionError>;
    async fn save(&self, session: &Session) -> Result<(), SessionError>;
    async fn delete(&self, email: &str) -> Result<(), SessionError>;
}

// -- In-memory store ----------------------------------------------------------

/// In-memory session store (not persisted across restarts).
pub struct MemorySessionStore {
    sessions: Mutex<HashMap<String, Session>>,
}

impl MemorySessionStore {
    pub fn new() -> Self {
        Self {
            sessions: Mutex::new(HashMap::new()),
        }
    }
}

impl Default for MemorySessionStore {
    fn default() -> Self {
        Self::new()
    }
}

#[async_trait]
impl SessionStore for MemorySessionStore {
    async fn load(&self, email: &str) -> Result<Session, SessionError> {
        let sessions = self.sessions.lock().await;
        sessions
            .get(email)
            .cloned()
            .ok_or(SessionError::NoSession)
    }

    async fn save(&self, session: &Session) -> Result<(), SessionError> {
        if session.email.is_empty() {
            return Err(SessionError::EmailRequired);
        }
        let mut s = session.clone();
        if s.updated_at == DateTime::<Utc>::default() {
            s.updated_at = Utc::now();
        }
        let mut sessions = self.sessions.lock().await;
        sessions.insert(s.email.clone(), s);
        Ok(())
    }

    async fn delete(&self, email: &str) -> Result<(), SessionError> {
        let mut sessions = self.sessions.lock().await;
        sessions.remove(email);
        Ok(())
    }
}

// -- JSON file store ----------------------------------------------------------

#[derive(Serialize, Deserialize)]
struct FileStoreData {
    sessions: HashMap<String, Session>,
}

/// JSON file-backed session store.
pub struct FileSessionStore {
    path: PathBuf,
    mu: Mutex<()>,
}

impl FileSessionStore {
    pub async fn new(path: impl AsRef<Path>) -> Result<Self, SessionError> {
        let path = path.as_ref().to_path_buf();
        if let Some(dir) = path.parent() {
            if dir != Path::new("") && dir != Path::new(".") {
                tokio::fs::create_dir_all(dir).await?;
            }
        }
        Ok(Self {
            path,
            mu: Mutex::new(()),
        })
    }

    async fn load_all(&self) -> Result<FileStoreData, SessionError> {
        match tokio::fs::read(&self.path).await {
            Ok(data) => {
                let d: FileStoreData =
                    serde_json::from_slice(&data).map_err(|e| SessionError::Decode {
                        path: self.path.display().to_string(),
                        source: e,
                    })?;
                Ok(d)
            }
            Err(e) if e.kind() == std::io::ErrorKind::NotFound => Ok(FileStoreData {
                sessions: HashMap::new(),
            }),
            Err(e) => Err(SessionError::Io(e)),
        }
    }

    async fn save_all(&self, data: &FileStoreData) -> Result<(), SessionError> {
        let bytes = serde_json::to_vec_pretty(data)?;
        let tmp = format!("{}.tmp", self.path.display());

        #[cfg(unix)]
        {
            use std::os::unix::fs::OpenOptionsExt;
            use tokio::io::AsyncWriteExt;
            let mut f = tokio::fs::OpenOptions::new()
                .write(true)
                .create(true)
                .truncate(true)
                .mode(0o600)
                .open(&tmp)
                .await
                .map_err(SessionError::Io)?;
            f.write_all(&bytes).await.map_err(SessionError::Io)?;
        }
        #[cfg(not(unix))]
        {
            tokio::fs::write(&tmp, &bytes)
                .await
                .map_err(SessionError::Io)?;
        }

        if let Err(e) = tokio::fs::rename(&tmp, &self.path).await {
            let _ = tokio::fs::remove_file(&tmp).await;
            return Err(SessionError::Io(e));
        }
        Ok(())
    }
}

#[async_trait]
impl SessionStore for FileSessionStore {
    async fn load(&self, email: &str) -> Result<Session, SessionError> {
        let _guard = self.mu.lock().await;
        let data = self.load_all().await?;
        let mut session = data
            .sessions
            .get(email)
            .cloned()
            .ok_or(SessionError::NoSession)?;
        if session.email.is_empty() {
            session.email = email.to_string();
        }
        Ok(session)
    }

    async fn save(&self, session: &Session) -> Result<(), SessionError> {
        if session.email.is_empty() {
            return Err(SessionError::EmailRequired);
        }
        let _guard = self.mu.lock().await;
        let mut data = self.load_all().await?;
        let mut stored = session.clone();
        if stored.updated_at == DateTime::<Utc>::default() {
            stored.updated_at = Utc::now();
        }
        data.sessions.insert(stored.email.clone(), stored);
        self.save_all(&data).await
    }

    async fn delete(&self, email: &str) -> Result<(), SessionError> {
        let _guard = self.mu.lock().await;
        let mut data = self.load_all().await?;
        if !data.sessions.contains_key(email) {
            return Ok(());
        }
        data.sessions.remove(email);
        self.save_all(&data).await
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[tokio::test]
    async fn memory_store_round_trip() {
        let store = MemorySessionStore::new();
        let session = Session {
            email: "test@example.com".into(),
            user_id: 42,
            access_token: "access".into(),
            refresh_token: "refresh".into(),
            device_uuid: "00000000-0000-0000-0000-000000000000".into(),
            oauth_key: String::new(),
            updated_at: Utc::now(),
        };

        store.save(&session).await.unwrap();
        let loaded = store.load("test@example.com").await.unwrap();
        assert_eq!(loaded.user_id, 42);
        assert_eq!(loaded.access_token, "access");
        assert_eq!(loaded.refresh_token, "refresh");

        store.delete("test@example.com").await.unwrap();
        assert!(matches!(
            store.load("test@example.com").await,
            Err(SessionError::NoSession)
        ));
    }

    #[tokio::test]
    async fn memory_store_no_session() {
        let store = MemorySessionStore::new();
        assert!(matches!(
            store.load("nobody@example.com").await,
            Err(SessionError::NoSession)
        ));
    }

    #[tokio::test]
    async fn memory_store_email_required() {
        let store = MemorySessionStore::new();
        let session = Session {
            email: String::new(),
            user_id: 1,
            access_token: "a".into(),
            refresh_token: "r".into(),
            device_uuid: String::new(),
            oauth_key: String::new(),
            updated_at: Utc::now(),
        };
        assert!(matches!(
            store.save(&session).await,
            Err(SessionError::EmailRequired)
        ));
    }

    #[tokio::test]
    async fn memory_store_delete_missing_is_ok() {
        let store = MemorySessionStore::new();
        store.delete("nonexistent@example.com").await.unwrap();
    }

    #[tokio::test]
    async fn file_store_round_trip() {
        let dir = tempfile::tempdir().unwrap();
        let path = dir.path().join("sessions.json");
        let store = FileSessionStore::new(&path).await.unwrap();

        let session = Session {
            email: "test@example.com".into(),
            user_id: 99,
            access_token: "at".into(),
            refresh_token: "rt".into(),
            device_uuid: "dev-uuid".into(),
            oauth_key: String::new(),
            updated_at: Utc::now(),
        };

        store.save(&session).await.unwrap();
        let loaded = store.load("test@example.com").await.unwrap();
        assert_eq!(loaded.user_id, 99);
        assert_eq!(loaded.device_uuid, "dev-uuid");

        // Verify the file is valid JSON with the expected schema.
        let raw = tokio::fs::read_to_string(&path).await.unwrap();
        let parsed: serde_json::Value = serde_json::from_str(&raw).unwrap();
        assert!(parsed["sessions"]["test@example.com"]["user_id"]
            .as_i64()
            .is_some());

        store.delete("test@example.com").await.unwrap();
        assert!(matches!(
            store.load("test@example.com").await,
            Err(SessionError::NoSession)
        ));
    }
}
