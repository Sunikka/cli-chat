package utils

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Username  string
	HashedPw  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
