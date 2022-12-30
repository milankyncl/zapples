package storage

import "time"

type Feature struct {
	Id          int
	Key         string
	Description *string
	Enabled     bool
	CreatedAt   time.Time
}
