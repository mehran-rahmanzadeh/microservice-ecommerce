package repository

import (
	"context"
	domain "github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/base"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/model"
	"gorm.io/gorm"
)

type categoryRepository struct {
	DB *gorm.DB
}

// FetchCategories will select parent categories and their children fetched
func (r *categoryRepository) FetchCategories(ctx context.Context) ([]model.Category, error) {
	var categories []model.Category
	var maxDepth int
	maxDepthQuery := "select max(depth) from categories"
	r.DB.Raw(maxDepthQuery).Scan(&maxDepth)
	preloadQuery := "Children"
	for i := 0; i < maxDepth-1; i += 1 {
		preloadQuery += ".Children"
	}
	tx := r.DB.Where("parent_id is null").Preload(preloadQuery).Find(&categories)
	return categories, tx.Error
}

func (r *categoryRepository) GetCategoryByID(ctx context.Context, id uint) (model.Category, error) {
	var category model.Category
	tx := r.DB.Find(&category, "id = ?", id)
	return category, tx.Error
}

func (r *categoryRepository) GetChildrenCategories(ctx context.Context, parentID uint) ([]model.Category, error) {
	var categories []model.Category
	var maxDepth int
	maxDepthQuery := "select max(depth) from categories"
	r.DB.Raw(maxDepthQuery).Scan(&maxDepth)
	preloadQuery := "Children"
	for i := 0; i < maxDepth-1; i += 1 {
		preloadQuery += ".Children"
	}
	tx := r.DB.Where("parent_id = ?", parentID).Preload(preloadQuery).Find(&categories)
	return categories, tx.Error
}

func (r *categoryRepository) GetParentCategories(ctx context.Context) ([]model.Category, error) {
	var categories []model.Category
	tx := r.DB.Where("parent_id is null").Find(&categories)
	return categories, tx.Error
}

func (r *categoryRepository) SearchCategories(ctx context.Context, query string) ([]model.Category, error) {
	var categories []model.Category
	tx := r.DB.Where("title LIKE ?", "%"+query+"%").Or("description LIKE ?", "%"+query+"%").Find(&categories)
	return categories, tx.Error
}

func NewCategoryRepository(db *gorm.DB) domain.CategoryRepository {
	return &categoryRepository{db}
}
