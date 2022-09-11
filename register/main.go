package main

import (
	"fmt"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/delivery/server"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/model"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/proto"
	repository "github.com/mehran-rahmanzadeh/microservice-ecommerce/repository/db"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/usecase"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
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
	err = db.AutoMigrate(model.User{})
	if err != nil {
		panic("failed to migrate User")
	}

	// init repo & usecase
	registerRepo := repository.NewRegisterRepository(db)
	registerUsecase := usecase.NewRegisterUsecase(registerRepo)

	// init grpc
	lis, err := net.Listen("tcp", fmt.Sprintf(
		"%s:%d", viper.GetString("server.host"), viper.GetInt("server.port")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	proto.RegisterUserRegisterServer(grpcServer, server.NewRegisterUserServer(registerUsecase))
	_ = grpcServer.Serve(lis)
}
