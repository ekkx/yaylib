package yaylib

import (
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)

// Session is one row of the local session cache.
type Session struct {
	Email        string
	UserID       int64
	AccessToken  string
	RefreshToken string
	DeviceUUID   string
	UpdatedAt    time.Time
}

// openSessionStore opens (or creates) a SQLite database and ensures the
// `sessions` table exists. WAL mode keeps concurrent readers safe.
func openSessionStore(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path+"?_pragma=journal_mode(WAL)&_pragma=foreign_keys(1)")
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS sessions (
			email         TEXT PRIMARY KEY,
			user_id       INTEGER NOT NULL,
			access_token  TEXT NOT NULL,
			refresh_token TEXT NOT NULL,
			device_uuid   TEXT NOT NULL,
			updated_at    INTEGER NOT NULL
		)`); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

// LoadSession looks up a cached session by email. On success, tokens and
// DeviceUUID are applied to the Client so subsequent calls reuse them.
// Returns sql.ErrNoRows if no row matches `email`.
func (c *Client) LoadSession(email string) (*Session, error) {
	if c.sessionDB == nil {
		return nil, fmt.Errorf("yaylib: session store not configured (use WithSessionStore)")
	}
	row := c.sessionDB.QueryRow(
		`SELECT user_id, access_token, refresh_token, device_uuid, updated_at
		 FROM sessions WHERE email = ?`, email)
	s := &Session{Email: email}
	var updated int64
	if err := row.Scan(&s.UserID, &s.AccessToken, &s.RefreshToken, &s.DeviceUUID, &updated); err != nil {
		return nil, err
	}
	s.UpdatedAt = time.Unix(updated, 0)

	c.SetTokens(s.AccessToken, s.RefreshToken)
	if s.DeviceUUID != "" {
		c.DeviceUUID = s.DeviceUUID
	}
	return s, nil
}

// SaveSession writes (upserts) the given session and also activates its
// tokens on the Client. If UpdatedAt is zero it defaults to time.Now().
func (c *Client) SaveSession(s *Session) error {
	if c.sessionDB == nil {
		return fmt.Errorf("yaylib: session store not configured (use WithSessionStore)")
	}
	if s.UpdatedAt.IsZero() {
		s.UpdatedAt = time.Now()
	}
	if s.DeviceUUID == "" {
		s.DeviceUUID = c.DeviceUUID
	}
	_, err := c.sessionDB.Exec(`
		INSERT INTO sessions (email, user_id, access_token, refresh_token, device_uuid, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
		ON CONFLICT(email) DO UPDATE SET
			user_id       = excluded.user_id,
			access_token  = excluded.access_token,
			refresh_token = excluded.refresh_token,
			device_uuid   = excluded.device_uuid,
			updated_at    = excluded.updated_at`,
		s.Email, s.UserID, s.AccessToken, s.RefreshToken, s.DeviceUUID, s.UpdatedAt.Unix())
	if err != nil {
		return err
	}
	c.SetTokens(s.AccessToken, s.RefreshToken)
	if s.DeviceUUID != "" {
		c.DeviceUUID = s.DeviceUUID
	}
	return nil
}

// DeleteSession removes the cached row for `email`, if any.
func (c *Client) DeleteSession(email string) error {
	if c.sessionDB == nil {
		return fmt.Errorf("yaylib: session store not configured (use WithSessionStore)")
	}
	_, err := c.sessionDB.Exec(`DELETE FROM sessions WHERE email = ?`, email)
	return err
}
