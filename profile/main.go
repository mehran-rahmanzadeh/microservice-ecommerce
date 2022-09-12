package main

import (
	"fmt"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/delivery/server"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/proto"
	repository "github.com/mehran-rahmanzadeh/microservice-ecommerce/repository/grpc"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/usecase"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
)

func main()  {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// init clients
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	registerConn, err := grpc.Dial(viper.GetString("services.register.url"), opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer registerConn.Close()
	registerClient := proto.NewUserRegisterClient(registerConn)
	loginConn, err := grpc.Dial(viper.GetString("services.login.url"), opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer loginConn.Close()
	loginClient := proto.NewUserLoginClient(loginConn)

	// init repo & usecase
	profileRepo := repository.NewProfileRepository(registerClient)
	profileUsecase := usecase.NewProfileUsecase(profileRepo, loginClient)

	// init grpc
	lis, err := net.Listen("tcp", fmt.Sprintf(
		"%s:%d", viper.GetString("server.host"), viper.GetInt("server.port")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var serverOpts []grpc.ServerOption
	grpcServer := grpc.NewServer(serverOpts...)
	proto.RegisterUserProfileServer(grpcServer, server.NewUserProfileServer(profileUsecase))
	_ = grpcServer.Serve(lis)
}
