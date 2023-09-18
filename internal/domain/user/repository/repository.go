package repository

import (
	"context"
	"fmt"
	"memo/ent"
	entUser "memo/ent/user"
	"memo/internal/domain/user"
	"memo/internal/domain/user/entity"

	"golang.org/x/crypto/bcrypt"
)

type Repository struct {
	ent *ent.Client
}

func New(ent *ent.Client) *Repository {
	return &Repository{
		ent: ent,
	}
}

func (r *Repository) Register(ctx context.Context, req *user.RegisterRequest) (*entity.User, error) {
	user, _ := r.ent.User.Query().Where(entUser.Name(req.Name)).Only(ctx)
	if user != nil {
		return nil, fmt.Errorf("That username is already in use!")
	}
	if req.Password != req.PasswordConfirm {
		return nil, fmt.Errorf("Passwords do not match!")
	}
	hashHassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	user, err := r.ent.User.
		Create().
		SetName(req.Name).
		SetPassword(hashHassword).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %v", err)
	}

	resp := &entity.User{
		ID:       user.ID,
		Name:     user.Name,
		CretedAt: user.CreatedAt,
	}
	return resp, nil
}

func (r *Repository) List(ctx context.Context) ([]*entity.User, error) {
	users, err := r.ent.User.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed query user: %v", err)
	}
	resp := make([]*entity.User, 0)
	for _, v := range users {
		resp = append(resp, &entity.User{
			ID:       v.ID,
			Name:     v.Name,
			CretedAt: v.CreatedAt,
		})
	}
	return resp, nil
}

func (r *Repository) Login(ctx context.Context, req *user.LoginRequest) (*entity.User, error) {
	user, err := r.ent.User.Query().Where(entUser.Name(req.Name)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("User not found")
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(req.Password)); err != nil {
		return nil, fmt.Errorf("Login failed: wrong password")
	}

	resp := &entity.User{
		ID:       user.ID,
		Name:     user.Name,
		CretedAt: user.CreatedAt,
	}
	return resp, nil
}
