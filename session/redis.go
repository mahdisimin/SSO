package session

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client *redis.Client
	Ctx    context.Context
}

/*
	Create(session Session) error
	Get(ID string) (Session, error)
	Delete(ID string) error
*/

var ErrExpired = errors.New("expired session")

func NewRedis(client *redis.Client) *Redis {
	return &Redis{client, context.Background()}
}

func (r *Redis) Create(session Session) error {
	sessionJson, errMarshal := json.Marshal(session)
	if errMarshal != nil {
		return errMarshal
	}
	if time.Until(session.ExpiresAt) < 0 {
		return ErrExpired
	}
	if err := r.Client.Set(r.Ctx, session.ID, sessionJson, time.Until(session.ExpiresAt)).Err(); err != nil {
		return err
	}
	return nil

}
func (r *Redis) Get(sessionID string) (Session, error) {
	sessionJsn, errGet := r.Client.Get(r.Ctx, sessionID).Result()
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
	if err := r.Client.Del(r.Ctx, ID).Err(); err != nil {
		return err
	}
	return nil
}
