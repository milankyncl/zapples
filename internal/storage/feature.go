package storage

import "time"

type Feature struct {
	Id           int
	Key          string
	Description  *string
	Enabled      bool
	EnabledSince *time.Time
	EnabledUntil *time.Time
	CreatedAt    time.Time
}
