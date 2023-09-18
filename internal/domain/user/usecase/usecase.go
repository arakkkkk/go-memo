package usecase

import (
	"context"
	"fmt"
	"memo/internal/domain/user"
	"memo/internal/domain/user/entity"
	"memo/internal/domain/user/repository"
	authUtil "memo/internal/util/auth"
)

type Usecase struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) Register(ctx context.Context, req *user.RegisterRequest) (string, error) {
	user, err := u.repo.Register(ctx, req)
	if err != nil {
		return "", fmt.Errorf("%s", err)
	}
	token, err := authUtil.GenJwt(user.ID)
	if err != nil {
		return "", fmt.Errorf("%s", err)
	}
	return token, nil
}

func (u *Usecase) Login(ctx context.Context, req *user.LoginRequest) (string, error) {
	user, err := u.repo.Login(ctx, req)
	if err != nil {
		return "", fmt.Errorf("%s", err)
	}
	token, err := authUtil.GenJwt(user.ID)
	if err != nil {
		return "", fmt.Errorf("%s", err)
	}
	return token, nil
}

func (u *Usecase) List(ctx context.Context) ([]*entity.User, error) {
	return u.repo.List(ctx)
}
