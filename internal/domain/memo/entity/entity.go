package entity

import "time"

type Memo struct {
	UserID      string
	Description string
	CretedAt    time.Time
}
