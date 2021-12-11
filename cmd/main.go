package main

import (
	"github.com/KonstantinP85/grpc-server/pkg/api"
	"github.com/KonstantinP85/grpc-server/pkg/news"
	"github.com/KonstantinP85/grpc-server/pkg/repository"
	"github.com/KonstantinP85/grpc-server/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error init config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error env var: %s", err.Error())
	}

	db, err := repository.NewDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	repo := repository.NewRepository(db)
	serv := service.NewService(repo)

	s := grpc.NewServer()
	srv := news.NewServer(serv)
	api.RegisterNewsServiceServer(s, srv)

	l, err := net.Listen("tcp", ":8010")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
