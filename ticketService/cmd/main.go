package main

import (
	"fmt"
	"google.golang.org/grpc"
	"intelXlabs/ticketService/pkg"
	"intelXlabs/ticketService/proto"
	"log"
	"net"
)

func main() {
	log.Println("starting ticket service at port 8005")
	lis, err := net.Listen("tcp", ":8005")
	if err != nil {
		fmt.Println("failed to listen:", err)
		return
	}
	s := grpc.NewServer()
	proto.RegisterTicketServiceServer(s, &pkg.TicketService{})
	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to start server :", err)
		return
	}
}
