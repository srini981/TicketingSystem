package pkg

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	rbac "intelXlabs/rbacService/proto"
	ticket "intelXlabs/ticketService/proto"
	user "intelXlabs/userService/proto"
	"log"
)

// client struct to handle all the grpc connections
type Client struct {
	RbacClient   rbac.RbacServiceClient
	RbacConn     *grpc.ClientConn
	TicketClient ticket.TicketServiceClient
	TicketConn   *grpc.ClientConn
	UserClient   user.UserServiceClient
	UserConn     *grpc.ClientConn
}

var client = InitClients()

// initclient is used to initilize all the grpc clients and connections
func InitClients() Client {
	client := Client{}
	rbacClient, rbacConn, err := RbacClient(":8006")
	if err != nil {
		log.Fatalf("could't connect to rbac service err: %v", err)
		return Client{}
	}
	client.RbacClient = rbacClient
	client.RbacConn = rbacConn
	ticketClient, ticketConn, err := TicketClient(":8008")
	if err != nil {
		log.Fatalf("could't connect to ticket service err: %v", err)
		return Client{}
	}
	client.TicketClient = ticketClient
	client.TicketConn = ticketConn
	userClient, userConn, err := UserClient(":8003")
	if err != nil {
		log.Fatalf("could't connect to user service err: %v", err)
		return Client{}
	}

	client.UserClient = userClient
	client.UserConn = userConn

	return client
}

func RbacClient(port string) (rbac.RbacServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}
	rbacClient := rbac.NewRbacServiceClient(conn)
	return rbacClient, conn, nil
}

func TicketClient(port string) (ticket.TicketServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}
	ticketClient := ticket.NewTicketServiceClient(conn)
	return ticketClient, conn, nil
}

func UserClient(port string) (user.UserServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}
	userClient := user.NewUserServiceClient(conn)
	return userClient, conn, nil
}
