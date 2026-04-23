package yaylib

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Session is one persisted login record.
type Session struct {
	Email        string    `json:"email,omitempty"`
	UserID       int64     `json:"user_id"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	DeviceUUID   string    `json:"device_uuid,omitempty"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// ErrNoSession is returned by SessionStore.Load when no record exists for
// the given email.
var ErrNoSession = errors.New("yaylib: no session for that email")

// SessionStore persists (email → Session) pairs so tokens survive across
// process restarts. Implementations must be safe for concurrent use from
// within a single process. Multi-process writers are out of scope.
type SessionStore interface {
	// Load returns the session for email, or ErrNoSession if none is stored.
	Load(email string) (*Session, error)

	// Save upserts the session keyed by s.Email.
	Save(s *Session) error

	// Delete removes the session for email. A missing key is not an error.
	Delete(email string) error
}

// -- in-memory store --------------------------------------------------------

type memoryStore struct {
	mu       sync.Mutex
	sessions map[string]Session
}

// NewMemoryStore returns a SessionStore that lives only in process memory.
// Useful for tests and short-lived scripts.
func NewMemoryStore() SessionStore {
	return &memoryStore{sessions: map[string]Session{}}
}

func (m *memoryStore) Load(email string) (*Session, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	s, ok := m.sessions[email]
	if !ok {
		return nil, ErrNoSession
	}
	out := s
	return &out, nil
}

func (m *memoryStore) Save(s *Session) error {
	if s.Email == "" {
		return fmt.Errorf("yaylib: SessionStore.Save: Email is required")
	}
	if s.UpdatedAt.IsZero() {
		s.UpdatedAt = time.Now()
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	m.sessions[s.Email] = *s
	return nil
}

func (m *memoryStore) Delete(email string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.sessions, email)
	return nil
}

// -- JSON file store --------------------------------------------------------

type fileStore struct {
	path string
	mu   sync.Mutex
}

type fileStoreData struct {
	Sessions map[string]*Session `json:"sessions"`
}

// NewSessionStore returns a SessionStore backed by a JSON file at path.
// This is the default implementation; use NewMemoryStore for an in-process
// alternative. The parent directory is created if missing; the file itself
// is created lazily on the first Save. Writes are atomic within a single
// process via temp-file + rename. Concurrent writers across multiple
// processes are not supported — use one session file per process.
func NewSessionStore(path string) (SessionStore, error) {
	if dir := filepath.Dir(path); dir != "" && dir != "." {
		if err := os.MkdirAll(dir, 0o700); err != nil {
			return nil, fmt.Errorf("yaylib: create session dir: %w", err)
		}
	}
	return &fileStore{path: path}, nil
}

func (f *fileStore) loadAll() (*fileStoreData, error) {
	data, err := os.ReadFile(f.path)
	if errors.Is(err, os.ErrNotExist) {
		return &fileStoreData{Sessions: map[string]*Session{}}, nil
	}
	if err != nil {
		return nil, err
	}
	var d fileStoreData
	if err := json.Unmarshal(data, &d); err != nil {
		return nil, fmt.Errorf("yaylib: decode %s: %w", f.path, err)
	}
	if d.Sessions == nil {
		d.Sessions = map[string]*Session{}
	}
	return &d, nil
}

func (f *fileStore) saveAll(d *fileStoreData) error {
	b, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return err
	}
	tmp := f.path + ".tmp"
	if err := os.WriteFile(tmp, b, 0o600); err != nil {
		return err
	}
	if err := os.Rename(tmp, f.path); err != nil {
		os.Remove(tmp)
		return err
	}
	return nil
}

func (f *fileStore) Load(email string) (*Session, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	d, err := f.loadAll()
	if err != nil {
		return nil, err
	}
	s, ok := d.Sessions[email]
	if !ok {
		return nil, ErrNoSession
	}
	if s.Email == "" {
		s.Email = email
	}
	return s, nil
}

func (f *fileStore) Save(s *Session) error {
	if s.Email == "" {
		return fmt.Errorf("yaylib: SessionStore.Save: Email is required")
	}
	f.mu.Lock()
	defer f.mu.Unlock()
	d, err := f.loadAll()
	if err != nil {
		return err
	}
	if s.UpdatedAt.IsZero() {
		s.UpdatedAt = time.Now()
	}
	stored := *s
	d.Sessions[s.Email] = &stored
	return f.saveAll(d)
}

func (f *fileStore) Delete(email string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	d, err := f.loadAll()
	if err != nil {
		return err
	}
	if _, ok := d.Sessions[email]; !ok {
		return nil
	}
	delete(d.Sessions, email)
	return f.saveAll(d)
}

// -- Convenience methods on *Client -----------------------------------------

// LoadSession loads the cached session for email and applies its tokens and
// DeviceUUID to the client. Returns ErrNoSession when nothing is cached.
// Errors when no session store is configured.
func (c *Client) LoadSession(email string) (*Session, error) {
	if c.sessionStore == nil {
		return nil, fmt.Errorf("yaylib: session store not configured (use WithSessionStore)")
	}
	s, err := c.sessionStore.Load(email)
	if err != nil {
		return nil, err
	}
	c.SetTokens(s.AccessToken, s.RefreshToken)
	if s.DeviceUUID != "" {
		c.DeviceUUID = s.DeviceUUID
	}
	return s, nil
}

// SaveSession writes the session and also activates its tokens on the
// client. Errors when no session store is configured.
func (c *Client) SaveSession(s *Session) error {
	if c.sessionStore == nil {
		return fmt.Errorf("yaylib: session store not configured (use WithSessionStore)")
	}
	if s.UpdatedAt.IsZero() {
		s.UpdatedAt = time.Now()
	}
	if s.DeviceUUID == "" {
		s.DeviceUUID = c.DeviceUUID
	}
	if err := c.sessionStore.Save(s); err != nil {
		return err
	}
	c.SetTokens(s.AccessToken, s.RefreshToken)
	if s.DeviceUUID != "" {
		c.DeviceUUID = s.DeviceUUID
	}
	return nil
}

// DeleteSession removes the cached session for email.
func (c *Client) DeleteSession(email string) error {
	if c.sessionStore == nil {
		return fmt.Errorf("yaylib: session store not configured (use WithSessionStore)")
	}
	return c.sessionStore.Delete(email)
}
