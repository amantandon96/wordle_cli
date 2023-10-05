package models

import "time"

type Game struct {
	Id          int
	Word        string
	Mode        string
	State       string
	MaxAttempts int
	Attempts    []Attempt
	CreatedAt   time.Time
}

const (
	GameCreated               = "created"
	GameInProgress            = "in_progress"
	GameCompletedSuccessfully = "successful"
	GameFailed                = "failed"
	GameForfeited             = "forfeited"
)
