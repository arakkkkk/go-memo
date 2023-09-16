package usecase

import (
	"context"
	"memo/internal/domain/user"
	"memo/internal/domain/user/entity"
	"memo/internal/domain/user/repository"
)

type Usecase struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) Register(ctx context.Context, req *user.RegisterRequest) (*entity.User, error) {
	return u.repo.Register(ctx, req)
}

func (u *Usecase) Login(ctx context.Context, req *user.LoginRequest) (string, error) {
	return u.repo.Login(ctx, req)
}

func (u *Usecase) List(ctx context.Context) ([]*entity.User, error) {
	return u.repo.List(ctx)
}
