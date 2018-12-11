package repository

import (
	"context"
	"web_blog/model"
)

type PostRepository interface {
	Fetch(ctx context.Context, num int64) ([]*model.Post, error)
	FindById(ctx context.Context, id int64) (*model.Post, error)
	Create(ctx context.Context, post *model.Post) (int64, error)
	Update(ctx context.Context, post *model.Post) (*model.Post, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
