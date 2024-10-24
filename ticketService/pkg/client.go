package pkg

import (
	"github.com/IBM/sarama"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"intelXlabs/dbService/proto"
	"log"
)

// client struct used to handle grpc clients and connections
type Client struct {
	DbServiceClient proto.DbServiceClient
	Dbconn          *grpc.ClientConn
	kafkaProducer   sarama.SyncProducer
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
	brokersUrl := []string{"localhost:29092"}
	kafkaProducer, err := KafkaProducer(brokersUrl)

	if err != nil {
		log.Fatalf("failed to create client", err.Error())
		return Client{}
	}

	client.kafkaProducer = kafkaProducer
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

func KafkaProducer(brokersUrl []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer(brokersUrl, config)

	if err != nil {
		return nil, err
	}

	return conn, nil
}
