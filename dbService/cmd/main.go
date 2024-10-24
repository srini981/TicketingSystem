package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"intelXlabs/dbService/pkg"
	"intelXlabs/dbService/proto"
	"log"
	"net"
)

func main() {
	godotenv.Load()
	port := ":8002"
	log.Println("starting db service at port 8002")

	lis, err := net.Listen("tcp", port)

	if err != nil {
		fmt.Println("failed to listen:", err)
		return
	}

	s := grpc.NewServer()
	proto.RegisterDbServiceServer(s, &pkg.DbService{})

	if err = s.Serve(lis); err != nil {
		fmt.Println("failed to start server :", err)
		return
	}
}
