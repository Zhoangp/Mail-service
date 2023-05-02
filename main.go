package main

import (
	"fmt"
	"github.com/Zhoangp/Mail-service/config"
	"github.com/Zhoangp/Mail-service/internal/delivery/http"
	usecase2 "github.com/Zhoangp/Mail-service/internal/usecase"
	"github.com/Zhoangp/Mail-service/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	env := os.Getenv("ENV")
	fileName := "config/config-local.yml"
	if env == "app" {
		fileName = "config/config-app.yml"
	}
	cf, err := config.LoadConfig(fileName)
	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	usecase := usecase2.NewMailUsecase(cf)
	handler := http.NewMailHandler(usecase)
	lis, err := net.Listen("tcp", ":"+cf.Service.Port)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Mail Svc on", cf.Service.Port)
	grpcServer := grpc.NewServer()
	pb.RegisterMailServiceServer(grpcServer, handler)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
