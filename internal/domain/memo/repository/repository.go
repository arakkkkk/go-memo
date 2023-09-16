package repository

import (
	"context"
	"fmt"
	"memo/ent"
	"memo/internal/domain/memo"
	"memo/internal/domain/memo/entity"
)

type Repository struct{
	ent *ent.Client
}

func New(ent *ent.Client) *Repository {
	return &Repository{
		ent: ent,
	}
}

func entBindSchema(m *ent.Memo) *entity.Memo {
	resp := &entity.Memo{
		Description: m.Description,
		CretedAt: m.CreatedAt,
	}
	return resp
}
func entMemosBindEntiryMemos(t []*ent.Memo) []*entity.Memo {
	resp := make([]*entity.Memo, 0)
	for _, v := range t {
		resp = append(resp, entBindSchema(v))
	}
	return resp
}

func (r *Repository) List(ctx context.Context) ([]*entity.Memo, error) {
  memos, err := r.ent.Memo.Query().All(ctx)
  resp := entMemosBindEntiryMemos(memos)
  if err != nil {
		return nil, fmt.Errorf("failed query user: %v", err)
	}
  return resp, nil
}

func (r *Repository) Create(ctx context.Context, req *memo.CreateRequest) (*entity.Memo, error) {
  memo, err := r.ent.Memo.
    Create().
    SetDescription(req.Description).
    Save(ctx)

  if err != nil {
		return nil, fmt.Errorf("failed creating memo: %v", err)
	}

	entityMemo := entBindSchema(memo)
  return entityMemo, nil
}

