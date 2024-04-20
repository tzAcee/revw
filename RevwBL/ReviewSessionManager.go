package RevwBL

import "revw/Logger"

type ReviewSessionManager struct {
	sessions map[string]ReviewSession
}

var globalSessionManager *ReviewSessionManager = nil

func newSessionsManager() *ReviewSessionManager {
	return &ReviewSessionManager{make(map[string]ReviewSession)}
}

func GetSessionsManager() *ReviewSessionManager {
	if globalSessionManager == nil {
		globalSessionManager = newSessionsManager()
	}

	return globalSessionManager
}

func (r *ReviewSessionManager) CreateReviewSession(text *string) *ReviewSession {
	newSession := NewReviewSession(text)
	_, ok := r.sessions[newSession.ID]
	if ok {
		Logger.GetLogger().Println("Trying to create a session with a duplicate ID...")
		return nil
	}

	r.sessions[newSession.ID] = *newSession

	return newSession
}

func (r *ReviewSessionManager) GetSessionById(id string) (*ReviewSession, bool) {
	session, ok := r.sessions[id]

	return &session, ok
}
