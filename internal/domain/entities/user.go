package entities

import "time"

type UserStatus string

const (
	Active   UserStatus = "active"
	Inactive UserStatus = "inactive"
)

type User struct {
	ID        int // TODO: Change to uuid.UUID
	Username  string
	Password  string
	Email     string
	Status    UserStatus
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
