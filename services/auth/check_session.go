package auth

import "SSO/user"

// CheckSession
//
//	comment the functionality
func (s *AuthService) CheckSession(sessionID string) (user.User, bool) {
	sessionIns, errSession := s.Sessions.Get(sessionID)
	if errSession != nil {
		return user.User{}, false
	}
	userID := sessionIns.UserID
	userIns, errFindUserByID := s.Users.FindByID(userID)
	if errFindUserByID != nil {
		return user.User{}, false
	}
	return userIns, true

}
