package usecase

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	domain "github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/base"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/model"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/proto"
	"github.com/spf13/viper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

var (
	jwtKey = viper.GetString("secrets.jwt_key")
)

type loginUsecase struct {
	repo domain.LoginRepository
}

func NewLoginUsecase(repo domain.LoginRepository) domain.LoginUsecase {
	return &loginUsecase{repo}
}

func (u loginUsecase) LoginUser(ctx context.Context, input *proto.LoginInput) (*proto.Token, error) {
	credentials := proto.ValidateCredentialsInput{
		Email:    input.Email,
		Password: input.Password,
	}
	isAuthenticated, err := u.repo.Authenticate(ctx, &credentials)
	if err != nil {
		return &proto.Token{}, status.Error(codes.PermissionDenied, err.Error())
	}
	if isAuthenticated {
		token, err := GenerateToken(input.Email)
		if err != nil {
			panic(err)
		}
		protoToken := &proto.Token{
			Access:  token,
			Refresh: "", // TODO: add refresh token
		}
		return protoToken, nil
	} else {
		return &proto.Token{}, status.Error(codes.PermissionDenied, "Incorrect input")
	}
}

func (u loginUsecase) AuthenticateToken(ctx context.Context, input *proto.TokenInput) (*proto.Authentication, error) {
	claims, err := ValidateToken(input.Access)
	if err != nil {
		return &proto.Authentication{}, status.Error(codes.PermissionDenied, err.Error())
	}
	return &proto.Authentication{
		IsAuthenticated: true,
		Email:              claims.Email,
	}, err
}

func GenerateToken(email string) (tokenString string, err error) {
	expirationTime := time.Now().Add(viper.GetDuration("secrets.jwt_lifetime") * time.Minute)
	claims:= &model.JWTClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(jwtKey))
	return
}

func ValidateToken(signedToken string) (claims *model.JWTClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&model.JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*model.JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}