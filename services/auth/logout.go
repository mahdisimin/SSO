package auth

func (a *AuthService) Logout(sessionID string) error {
	return a.Sessions.Delete(sessionID)
}
