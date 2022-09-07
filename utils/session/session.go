package session

import (
	"time"

	"github.com/google/uuid"
)

type client struct {
	UserId string
	Expiry time.Time
}

var sessions = make(map[string]*client)

//NewClient will return session.client , UserId should a Unique user identity
func NewClient(UserId string) *client {
	return &client{
		UserId: UserId,
		Expiry: time.Now().Add(360 * time.Second),
	}
}

func (c *client) CreateNewSession() (sessionToken string) {
	sessionToken = uuid.NewString()
	sessions[sessionToken] = c
	return sessionToken
}

// func CreateFreshSession(UserId string) (sessionToken string) {
// 	client := NewClient(UserId)
// 	sessionToken = uuid.NewString()
// 	sessions[sessionToken] = client
// 	return sessionToken
// }
func IsTokenValid(sessionToken string) (*client, bool) {
	client, ok := sessions[sessionToken]
	return client, ok
}
func (c *client) IsSessionExpired() bool {
	return c.Expiry.Before(time.Now())
}

func DestroySession(sessionToken string) {
	delete(sessions, sessionToken)
}
