package main

import (
	"google.golang.org/grpc"
	"intelXlabs/userService/pkg"
	"intelXlabs/userService/proto"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8006")
	if err != nil {
		log.Fatalf("failed to listen:", err)
		return
	}
	log.Println("starting user service at port 8006")
	s := grpc.NewServer()
	proto.RegisterUserServiceServer(s, &pkg.UserService{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to start server :", err)
		return
	}
}
