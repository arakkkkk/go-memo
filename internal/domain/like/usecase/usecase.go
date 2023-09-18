package usecase

import (
	"context"
	"memo/internal/domain/like"
	"memo/internal/domain/like/entity"
	"memo/internal/domain/like/repository"

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

func (u *Usecase) List(ctx context.Context, userID int) ([]*entity.Like, error) {
	return u.repo.List(ctx, userID)
}

func (u *Usecase) Count(ctx context.Context, memoID uuid.UUID) (int, error) {
	return u.repo.Count(ctx, memoID)
}

func (u *Usecase) IsLiked(ctx context.Context, memoID uuid.UUID, userID int) (bool, error) {
	return u.repo.IsLiked(ctx, memoID, userID)
}

func (u *Usecase) Add(ctx context.Context, req *like.AddRequest) (*entity.Like, error) {
	return u.repo.Add(ctx, req)
}

func (u *Usecase) Delete(ctx context.Context, req *like.DeleteRequest) (int, error) {
	return u.repo.Delete(ctx, req)
}
