package pkg

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"intelXlabs/dbService/proto"
	"log"
)

// client struct used to handle grpc clients and connections
type Client struct {
	DbServiceClient proto.DbServiceClient
	Dbconn          *grpc.ClientConn
}

var client = Init()

// init function is used to initilize the grpc clients
func Init() Client {
	client := Client{}
	dbClient, dbConn, err := dbClient(":8004")

	if err != nil {
		log.Fatalf("could't connect to rbac service err: %v", err)
		return Client{}
	}

	client.DbServiceClient = dbClient
	client.Dbconn = dbConn

	return client
}

func dbClient(port string) (proto.DbServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, nil, err
	}
	
	dbClient := proto.NewDbServiceClient(conn)
	return dbClient, conn, nil
}
