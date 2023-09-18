package entity

import (
	"time"

	"github.com/google/uuid"
)

type Like struct {
  UserID int
  MemoID uuid.UUID
  CretedAt time.Time
}

