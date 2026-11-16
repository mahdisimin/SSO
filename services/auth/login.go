package auth

import (
	"SSO/session"
	"SSO/user"
	"errors"
	"time"
)

type AuthService struct {
	Users      user.Store
	Sessions   session.Store
	sessionTTL time.Duration
}

func NewAuthService(users user.Store, sessions session.Store, sessionTTL time.Duration) AuthService {
	return AuthService{
		Users:      users,
		Sessions:   sessions,
		sessionTTL: sessionTTL,
	}
}

func (s *AuthService) Login(userName string, password string) (session.Session, error) {
	userIns, errFindUser := s.Users.FindByUsername(userName)
	if errFindUser != nil {
		if errors.Is(errFindUser, user.ErrInvalidCredentials) {
			return session.Session{}, user.ErrInvalidCredentials
		}
		if errors.Is(errFindUser, user.ErrNotFound) {
			return session.Session{}, user.ErrNotFound
		}
		return session.Session{}, errFindUser
	}

	if ok := userIns.CheckPassword(password); ok != true {
		return session.Session{}, user.ErrInvalidCredentials
	}

	sessionIns, errSession := session.NewSession(userIns.ID, s.sessionTTL)
	if errSession != nil {
		return session.Session{}, errSession
	}
	if err := s.Sessions.Create(sessionIns); err != nil {
		return session.Session{}, err
	}

	return sessionIns, nil
}
