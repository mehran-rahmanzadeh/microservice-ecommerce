package domain

import (
	"context"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/model"
)

type CategoryRepository interface {
	FetchCategories(ctx context.Context) ([]model.Category, error)
	GetCategoryByID(ctx context.Context, id uint) (model.Category, error)
	GetChildrenCategories(ctx context.Context, parentID uint) ([]model.Category, error)
	GetParentCategories(ctx context.Context) ([]model.Category, error)
	SearchCategories(ctx context.Context, query string) ([]model.Category, error)
}

type CategoryUsecase interface {
	// TODO
}
