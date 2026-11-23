package session

import (
	"errors"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestRedis_Create(t *testing.T) {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	},
	)
	t.Cleanup(func() { client.Close() })
	redisIns := NewRedis(client)

	if err := client.Ping(redisIns.Ctx).Err(); err != nil {
		t.Skip("Redis is not available : ", err)
	}
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
}

func TestRedis_Delete(t *testing.T) {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	},
	)
	t.Cleanup(func() { client.Close() })
	redisIns := NewRedis(client)

	if err := client.Ping(redisIns.Ctx).Err(); err != nil {
		t.Skip("Redis is not available : ", err)
	}

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
	if !errors.Is(errGet, ErrNotFound) {
		t.Error(errGet)
	}

}

func TestRedis_CreateExpire(t *testing.T) {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	},
	)
	t.Cleanup(func() { client.Close() })
	redisIns := NewRedis(client)

	if err := client.Ping(redisIns.Ctx).Err(); err != nil {
		t.Skip("Redis is not available : ", err)
	}

	sessionIns, errSessoin := NewSession("user-1", -1*time.Minute*5)
	if errSessoin != nil {
		t.Error(errSessoin)
	}
	if err := redisIns.Create(sessionIns); err != nil {
		if !errors.Is(err, ErrExpired) {
			t.Error("error should be ErrExpired")
		}
	}
	_, errGet := redisIns.Get(sessionIns.ID)
	if !errors.Is(errGet, ErrNotFound) {
		t.Error(errGet)

	}

	t.Cleanup(func() { client.Close() })
}
