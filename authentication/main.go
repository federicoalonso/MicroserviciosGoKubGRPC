package main

import (
	"flag"
	"fmt"
	"log"
	"microservicios/authentication/repository"
	"microservicios/authentication/service"
	"microservicios/db"
	"microservicios/pb"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var (
	local bool
	port  int
)

func init() {
	flag.IntVar(&port, "port", 9015, "authentication service port")
	flag.BoolVar(&local, "local", true, "run authentication service local")
	flag.Parse()
}

func main() {
	if local {
		err := godotenv.Load()
		if err != nil {
			log.Panic(err)
		}
	}

	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	usersRepository := repository.NewUsersRepository(conn)
	authService := service.NewAuthService(usersRepository)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, authService)

	log.Printf("Authentication service running on [::]:%d\n", port)
	grpcServer.Serve(lis)
}
