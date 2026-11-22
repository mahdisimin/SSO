package session

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	client *redis.Client
	ctx    context.Context
}

/*
	Create(session Session) error
	Get(ID string) (Session, error)
	Delete(ID string) error
*/

var ErrExpired = errors.New("expired session")

func NewRedis() *Redis {
	return &Redis{client: redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	},
	), ctx: context.Background()}
}

func (r *Redis) Create(session Session) error {
	sessionJson, errMarshal := json.Marshal(session)
	if errMarshal != nil {
		return errMarshal
	}
	if time.Until(session.ExpiresAt) < 0 {
		return ErrExpired
	}
	if err := r.client.Set(r.ctx, session.ID, sessionJson, time.Until(session.ExpiresAt)).Err(); err != nil {
		return err
	}
	return nil

}
func (r *Redis) Get(sessionID string) (Session, error) {
	sessionJsn, errGet := r.client.Get(r.ctx, sessionID).Result()
	if errGet != nil {
		if errors.Is(errGet, redis.Nil) {
			return Session{}, ErrNotFound
		}
		return Session{}, errGet
	}
	sessionIns := Session{}
	if err := json.Unmarshal([]byte(sessionJsn), &sessionIns); err != nil {
		return Session{}, err
	}
	return sessionIns, nil
}

func (r *Redis) Delete(ID string) error {
	if err := r.client.Del(r.ctx, ID).Err(); err != nil {
		return err
	}
	return nil
}
