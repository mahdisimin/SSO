package session

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestRedis_Create(t *testing.T) {

	t.Helper()

	redisIns := NewRedis()
	sessionIns, errSessoin := NewSession("user-1", time.Minute*5)
	if errSessoin != nil {
		t.Error(errSessoin)
	}
	if err := redisIns.Create(sessionIns); err != nil {
		t.Error(err)
	}
	sessionIns2, errGet := redisIns.Get(sessionIns.ID)
	if errGet != nil {
		t.Error(errGet)
	}
	if sessionIns2.UserID != "user-1" {
		t.Error("session ins id should be user-1")
	}
	println("userID is :", sessionIns.UserID)
}

func TestRedis_Delete(t *testing.T) {

	t.Helper()

	redisIns := NewRedis()
	sessionIns, errSessoin := NewSession("user-1", time.Minute*5)
	if errSessoin != nil {
		t.Error(errSessoin)
	}
	if err := redisIns.Create(sessionIns); err != nil {
		t.Error(err)
	}

	if err := redisIns.Delete(sessionIns.ID); err != nil {
		t.Error(err)
	}
	_, errGet := redisIns.Get(sessionIns.ID)
	println("error is : ", errGet)
	if !errors.Is(errGet, ErrNotFound) {
		t.Error(errGet)
	}
}

func TestRedis_CreateExpire(t *testing.T) {
	t.Helper()
	redisIns := NewRedis()
	sessionIns, errSessoin := NewSession("user-1", -1*time.Minute*5)
	if errSessoin != nil {
		t.Error(errSessoin)
	}
	if err := redisIns.Create(sessionIns); err != nil {
		fmt.Println("error is : ", errSessoin)
		if !errors.Is(err, ErrExpired) {
			t.Error("error should be ErrExpired")
		}
	}
	_, errGet := redisIns.Get(sessionIns.ID)
	if !errors.Is(errGet, ErrNotFound) {
		t.Error(errGet)
		println("error is : ", errGet)
	}
}
