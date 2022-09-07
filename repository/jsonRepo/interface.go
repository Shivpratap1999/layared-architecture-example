package jsonRepo

import (
	"time"
)

type CashStorer interface {
	StoreNewObject(key string, value interface{}, expiry time.Duration) error
	DeleteObject(key string) error
	IsObjectExist(key string) bool
}
