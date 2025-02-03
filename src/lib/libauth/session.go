package libauth

import (
	"errors"
	"sync"
)

type SessionStore interface {
	Set(sessionID, userID string) error
	Get(sessionID string) (string, error)
	Delete(sessionID string) error
}

type InMemorySessionStore struct {
	sessions map[string]string
	mu       sync.RWMutex
}

func NewInMemorySessionStore() *InMemorySessionStore {
	return &InMemorySessionStore{
		sessions: make(map[string]string),
	}
}

func (s *InMemorySessionStore) Set(sessionID, userID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.sessions[sessionID] = userID
	return nil
}

func (s *InMemorySessionStore) Get(sessionID string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	userID, ok := s.sessions[sessionID]
	if !ok {
		return "", errors.New("session not found")
	}
	return userID, nil
}

func (s *InMemorySessionStore) Delete(sessionID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.sessions, sessionID)
	return nil
}
