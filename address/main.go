package main

import (
	"fmt"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/delivery/server"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/model"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/proto"
	repository "github.com/mehran-rahmanzadeh/microservice-ecommerce/repository/db"
	repository2 "github.com/mehran-rahmanzadeh/microservice-ecommerce/repository/grpc"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/usecase"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net"
)

func main() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// init DB
	// TODO: change sqlite to PostgreSQL
	db, err := gorm.Open(sqlite.Open(viper.GetString("db.name")), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	err = db.AutoMigrate(model.Address{})
	if err != nil {
		panic("failed to migrate Address")
	}

	// init clients
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	profileConn, err := grpc.Dial(viper.GetString("services.profile.url"), opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer profileConn.Close()
	profileClient := proto.NewUserProfileClient(profileConn)

	// init repos
	addressDBRepository := repository.NewAddressDBRepository(db)
	addressGrpcRepository := repository2.NewAddressGrpcRepository(profileClient)
	addressUsecase := usecase.NewAddressUsecase(addressDBRepository, addressGrpcRepository)

	// init grpc server
	lis, err := net.Listen("tcp", fmt.Sprintf(
		"%s:%d", viper.GetString("server.host"), viper.GetInt("server.port")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var serverOpts []grpc.ServerOption
	grpcServer := grpc.NewServer(serverOpts...)
	proto.RegisterUserAddressServer(grpcServer, server.NewUserAddressServer(addressUsecase))
	_ = grpcServer.Serve(lis)
}
