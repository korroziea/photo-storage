package domain

import "time"

type User struct {
	FirstName string
	Email     string
	Password  string
	CreatedAt time.Time
}
