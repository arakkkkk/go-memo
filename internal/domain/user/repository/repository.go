package repository

import (
	"context"
	"fmt"
	"memo/ent"
	entUser "memo/ent/user"
	"memo/internal/domain/user"
	"memo/internal/domain/user/entity"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
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
			Name:     v.Name,
			CretedAt: v.CreatedAt,
		})
	}
	return resp, nil
}

func (r *Repository) Login(ctx context.Context, req *user.LoginRequest) (string, error) {
	user, err := r.ent.User.Query().Where(entUser.Name(req.Name)).Only(ctx)
	if err != nil {
		return "", fmt.Errorf("User not found")
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(req.Password)); err != nil {
		return "", fmt.Errorf("Login failed: wrong password")
	}

	// JWT
	claims := jwt.StandardClaims{
		Issuer:    strconv.Itoa(user.ID),            // stringに型変換
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 有効期限
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return token, nil
}
