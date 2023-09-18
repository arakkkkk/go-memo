package entity

import "time"

type User struct {
	ID       int
	Name     string
	CretedAt time.Time
}
