package pkg

import (
	cp "intelXlabs/chatService/proto"
	"intelXlabs/dbService/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"log"
)

// client struct to handle all the grpc clients
type Client struct {
	DbServiceClient proto.DbServiceClient
	Dbconn          *grpc.ClientConn
	ChatClient      cp.ChatServiceClient
	ChatConn        *grpc.ClientConn
}

// client variable which will be used to hold the client values inside pkg package
var client = Init()

// init function to initilize all the grpc clients
func Init() Client {
	client := Client{}
	dbClient, dbConn, err := dbClient(":8002")

	if err != nil {
		log.Fatalf("could't connect to rbac service err: %v", err)
		return Client{}
	}

	client.DbServiceClient = dbClient
	client.Dbconn = dbConn
	chatClient, chatConn, err := chatClient(":8008")

	if err != nil {
		log.Fatalf("could't connect to chat service err: %v", err)
		return Client{}
	}

	client.ChatClient = chatClient
	client.ChatConn = chatConn

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

func chatClient(port string) (cp.ChatServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, nil, err
	}

	chatClient := cp.NewChatServiceClient(conn)
	return chatClient, conn, nil
}
