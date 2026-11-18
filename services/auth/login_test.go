package auth

import (
	"SSO/session"
	"SSO/user"
	"errors"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestAuthService_CorrectPassword(t *testing.T) {
	userStore := user.NewMemoryStore()
	sessionStore := session.NewMemoryStore()
	userIns := user.NewUser("Mahdi", "test@gmail.com", "1234")
	if err := userStore.Create(userIns); err != nil {
		t.Fatal(err)
	}
	authSvc := NewAuthService(userStore, sessionStore, time.Hour)
	sessionIns, errLogin := authSvc.Login(userIns.UserName, "1234")
	if errLogin != nil {
		t.Fatal(errLogin)
	}
	if sessionIns.UserID != userIns.ID {
		log.Printf("userID : %s \n", userIns.ID)
		log.Printf("session.userID : %s", sessionIns.UserID)
		t.Errorf("userID is wrong")
	}

}

func TestAuthService_WrongPassword(t *testing.T) {
	userStore := user.NewMemoryStore()
	sessionStore := session.NewMemoryStore()
	userIns := user.NewUser("Mahdi", "test@gmail.com", "1234")
	if err := userStore.Create(userIns); err != nil {
		t.Fatal(err)
	}
	authSvc := NewAuthService(userStore, sessionStore, time.Hour)
	_, errLogin := authSvc.Login(userIns.UserName, "12345")
	if errLogin == nil {
		t.Errorf("expecting error but got nil")
	}

}

func TestAuthService_UserNotExists(t *testing.T) {
	userStore := user.NewMemoryStore()
	sessionStore := session.NewMemoryStore()
	userIns := user.NewUser("Mahdi", "test@gmail.com", "1234")
	if err := userStore.Create(userIns); err != nil {
		t.Fatal(err)
	}
	authSvc := NewAuthService(userStore, sessionStore, time.Hour)
	_, errLoginNotExists := authSvc.Login("otherUser", "1234")
	fmt.Println(errLoginNotExists)
	if !errors.Is(errLoginNotExists, user.ErrInvalidCredentials) {
		t.Errorf("expecting error but got other")
	}
	_, errLoginWrongPass := authSvc.Login("Mahdi", "12345")
	fmt.Println(errLoginWrongPass)
	if !errors.Is(errLoginWrongPass, user.ErrInvalidCredentials) {
		t.Errorf("expecting error but got other")
	}

}
