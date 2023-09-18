package repository

import (
	"context"
	"fmt"
	"memo/ent"
	entLike "memo/ent/likerecord"
	"memo/internal/domain/like"
	"memo/internal/domain/like/entity"

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

// func (r *Repository) Get(ctx context.Context, memoID uuid.UUID) (*entity.Memo, error) {
// 	memo, err := r.ent.Memo.Query().Where(entMemo.ID(memoID)).Only(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("%v", err)
// 	}
// 	resp := &entity.Memo{
// 		Description: memo.Description,
// 		CretedAt:    memo.CreatedAt,
// 	}
// 	return resp, nil
// }

func (r *Repository) List(ctx context.Context, userID int) ([]*entity.Like, error) {
	var entLikes []*ent.LikeRecord
	var err error
	if userID == 0 {
		entLikes, err = r.ent.LikeRecord.Query().All(ctx)
	} else {
		entLikes, err = r.ent.LikeRecord.Query().Where(entLike.UserID(userID)).All(ctx)
	}
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	resp := make([]*entity.Like, 0)
	for _, v := range entLikes {
		resp = append(resp, &entity.Like{
			MemoID:   v.MemoID,
			UserID:   v.UserID,
			CretedAt: v.CreatedAt,
		})
	}
	return resp, nil
}

func (r *Repository) Count(ctx context.Context, memoID uuid.UUID) (int, error) {
	entLikesCount, err := r.ent.LikeRecord.Query().Where(entLike.MemoID(memoID)).Count(ctx)
	if err != nil {
		return 0, fmt.Errorf("%v", err)
	}
	return entLikesCount, nil
}

func (r *Repository) IsLiked(ctx context.Context, memoID uuid.UUID, userID int) (bool, error) {
	entLike, err := r.ent.LikeRecord.Query().Where(entLike.MemoID(memoID)).Only(ctx)
	if err != nil {
		return false, fmt.Errorf("%v", err)
	}
	if entLike == nil {
		return false, nil
	} else {
		return true, nil
	}
}

func (r *Repository) Add(ctx context.Context, req *like.AddRequest) (*entity.Like, error) {
	entLike, err := r.ent.LikeRecord.
		Create().
		SetUserID(req.UserID).
		SetMemoID(req.MemoID).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating memo: %v", err)
	}

	resp := &entity.Like{
		MemoID:   entLike.MemoID,
		UserID:   entLike.UserID,
		CretedAt: entLike.CreatedAt,
	}
	return resp, nil
}

func (r *Repository) Delete(ctx context.Context, req *like.DeleteRequest) (int, error) {
	entLike, err := r.ent.LikeRecord.
		Delete().
		Where(entLike.MemoID(req.MemoID), entLike.UserID(req.UserID)).
		Exec(ctx)

	if err != nil {
		return 0, fmt.Errorf("failed creating memo: %v", err)
	}

	return entLike, nil
}
