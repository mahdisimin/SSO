package auth

// CheckSession
//
//	comment the functionality
func (s *AuthService) CheckSession(sessionID string) bool {
	sessionIns, errSession := s.Sessions.Get(sessionID)
	if errSession != nil {
		return false
	}
	userID := sessionIns.UserID
	_, errFindUserByID := s.Users.FindByID(userID)
	if errFindUserByID != nil {
		return false
	}
	return true

}
