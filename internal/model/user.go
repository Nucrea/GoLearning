package model

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	Id uuid.UUID
	Name string
	Email string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
	// Contact []Contact
}