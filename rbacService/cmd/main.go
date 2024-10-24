package main

import (
	"fmt"
	"google.golang.org/grpc"
	"intelXlabs/rbacService/pkg"
	"intelXlabs/rbacService/proto"
	"log"
	"net"
)

func main() {
	log.Println("Starting RBAC Service at port 8003")
	lis, err := net.Listen("tcp", ":8003")
	if err != nil {
		fmt.Println("failed to listen:", err)
		return
	}
	s := grpc.NewServer()
	proto.RegisterRbacServiceServer(s, &pkg.RbactService{})

	if err = s.Serve(lis); err != nil {
		fmt.Println("failed to start server :", err)
		return
	}
}
