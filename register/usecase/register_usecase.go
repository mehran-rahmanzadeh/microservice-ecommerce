package usecase

import (
	"context"
	domain "github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/base"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/proto"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RegisterUsecase struct {
	repo domain.RegisterRepository
}

func NewRegisterUsecase(repo domain.RegisterRepository) domain.RegisterUsecase {
	return &RegisterUsecase{repo}
}

func (u *RegisterUsecase) RegisterUser(ctx context.Context, input *proto.RegisterInput) (*proto.User, error)  {
	// hash input password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	if err != nil {
		panic(err)
	}
	// replace password with hashed one
	input.Password = string(hashedPassword)

	user, err := u.repo.Register(ctx, input)
	if err != nil {
		return &proto.User{}, status.Error(codes.InvalidArgument, err.Error())
	}
	protoUser := &proto.User{
		Id:        uint64(user.ID),
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	return protoUser, nil
}

func (u RegisterUsecase) AuthenticateUser(ctx context.Context, input *proto.ValidateCredentialsInput) (*proto.Validate, error) {
	isValid, err := u.repo.Authenticate(ctx, input)
	if err != nil {
		return &proto.Validate{}, status.Error(codes.InvalidArgument, err.Error())
	}
	protoValidate := &proto.Validate{IsValid: isValid}
	return protoValidate, nil
}
