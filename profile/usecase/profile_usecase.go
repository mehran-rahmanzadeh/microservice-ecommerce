package usecase

import (
	"context"
	domain "github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/base"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type profileUsecase struct {
	loginClient proto.UserLoginClient
	repo domain.ProfileRepository
}

func NewProfileUsecase(repo domain.ProfileRepository, client proto.UserLoginClient) domain.ProfileUsecase {
	return &profileUsecase{client, repo}
}

func (u profileUsecase) GetUserProfile(ctx context.Context, input *proto.ProfileInput) (*proto.Profile, error) {
	protoToken := &proto.TokenInput{Access: input.Token}
	auth, err := u.loginClient.Authenticate(ctx, protoToken)
	if err != nil {
		return &proto.Profile{}, status.Error(codes.Unauthenticated, err.Error())
	}
	if auth.IsAuthenticated {
		getUserInput := &proto.GetUserInfoInput{Email: auth.Email}
		user, err := u.repo.GetUser(ctx, getUserInput)
		if err != nil {
			return &proto.Profile{}, status.Error(codes.Internal, err.Error())
		}
		profile := &proto.Profile{
			Id:        user.Id,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}
		return profile, nil
	} else {
		return &proto.Profile{}, status.Error(codes.Unauthenticated, "Not authenticated")
	}
}
