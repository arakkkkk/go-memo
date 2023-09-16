package memo

import "time"

type CreateRequest struct {
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"create_at"`
}
