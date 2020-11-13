package entity

import (
	"time"

	"github.com/google/uuid"
)

type Author struct {
	ID        uuid.UUID
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
