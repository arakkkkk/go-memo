package like

import "github.com/google/uuid"

type AddRequest struct {
	MemoID uuid.UUID `json:"memo_id"`
	UserID int       `json:"user_id"`
}

type DeleteRequest struct {
	MemoID uuid.UUID `json:"memo_id"`
	UserID int       `json:"user_id"`
}
