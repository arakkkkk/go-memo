package repository

import (
	"context"
	"fmt"
	"memo/ent"
	entMemo "memo/ent/memo"
	"memo/internal/domain/memo"
	"memo/internal/domain/memo/entity"

	"github.com/google/uuid"
)

type Repository struct {
	ent *ent.Client
}

func New(ent *ent.Client) *Repository {
	return &Repository{
		ent: ent,
	}
}

func (r *Repository) Get(ctx context.Context, memoID uuid.UUID) (*entity.Memo, error) {
	memo, err := r.ent.Memo.Query().Where(entMemo.ID(memoID)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	resp := &entity.Memo{
		Description: memo.Description,
		CretedAt:    memo.CreatedAt,
	}
	return resp, nil
}

func (r *Repository) List(ctx context.Context) ([]*entity.Memo, error) {
	memos, err := r.ent.Memo.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	resp := make([]*entity.Memo, 0)
	for _, v := range memos {
		resp = append(resp, &entity.Memo{
			Description: v.Description,
			CretedAt:    v.CreatedAt,
		})
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

	resp := &entity.Memo{
		Description: memo.Description,
		CretedAt:    memo.CreatedAt,
	}
	return resp, nil
}
