package models

import "time"
type Session struct{
	UserEmailId string
	ExpirationTime time.Duration
}