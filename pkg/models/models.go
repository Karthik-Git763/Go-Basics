package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord = errors.New("models: no record matching found")
	// Add a new ErrInvalidCredentials error. We'll use this when a user provides an invalid email address or password.
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	// Add a new ErrDuplicateEmail error. We'll use this when a user tries to register with an email address that's already in use.
	ErrDuplicateEmail = errors.New("models: duplicate email")
)

type Snippet struct {
	ID int
	Title string
	Content string
	Created time.Time
	Expires time.Time
}

type User struct {
	ID int
	Name string
	Email string
	HashedPassword []byte
	Created time.Time
	Activate bool
}
