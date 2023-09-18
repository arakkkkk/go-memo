package usecase

import (
	"context"
	"memo/internal/domain/memo"
	"memo/internal/domain/memo/entity"
	"memo/internal/domain/memo/repository"

	"github.com/google/uuid"
)

type Usecase struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) Create(ctx context.Context, req *memo.CreateRequest) (*entity.Memo, error) {
	return u.repo.Create(ctx, req)
}

func (u *Usecase) List(ctx context.Context) ([]*entity.Memo, error) {
	return u.repo.List(ctx)
}

func (u *Usecase) Get(ctx context.Context, memoID uuid.UUID) (*entity.Memo, error) {
	return u.repo.Get(ctx, memoID)
}
