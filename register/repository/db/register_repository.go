package repository

import (
	"context"
	"errors"
	domain "github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/base"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/model"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/proto"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterRepository struct {
	DB *gorm.DB
}

// NewRegisterRepository initializes new repository for register
func NewRegisterRepository(db *gorm.DB) domain.RegisterRepository {
	return &RegisterRepository{db}
}


// Register new user to db (password in the input should be hashed)
func (r *RegisterRepository) Register(ctx context.Context, input *proto.RegisterInput) (model.User, error) {
	user := model.User{Email: input.Email, Password: input.Password}
	tx := r.DB.Create(&user)
	return user, tx.Error
}

// Authenticate user credentials and return true/false
func (r *RegisterRepository) Authenticate(ctx context.Context, input *proto.ValidateCredentialsInput) (bool, error) {
	var result model.User
	tx := r.DB.Model(model.User{Email: input.Email}).First(&result)
	if tx.Error != nil {
		return false, tx.Error
	}
	if result.Email != input.Email {
		return false, errors.New("user not found")
	}
	err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(input.Password))
	return err == nil, nil
}

func (r *RegisterRepository) GetUser(ctx context.Context, input *proto.GetUserInfoInput) (model.User, error) {
	var user model.User
	tx := r.DB.Where(&model.User{
		Email: input.Email,
	}).First(&user)
	if tx.Error != nil {
		return model.User{}, tx.Error
	}
	return user, nil
}
