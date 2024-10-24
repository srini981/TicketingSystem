package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/resend/resend-go/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// msg struct for handling kafka msg
type kafkaMsg struct {
	TicketID    int
	Description string
	Message     string
}

// consumer init function
func ConsumerInit() {
	log.Println("starting Consumer")
	topic := "tickets"
	worker, err := connectConsumer([]string{"localhost:29092"})
	if err != nil {
		log.Fatalf("couldt connect to kafka", err.Error())
		return
	}

	// Calling ConsumePartition. It will open one connection per broker
	// and share it for all partitions that live on it.
	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("couldt consume from topic in kafka", err.Error())
		return
	}

	log.Println("Consumer started ")
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	// Count how many message processed
	msgCount := 0

	// Get signal for finish
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				msgCount++
				fmt.Printf("Received message Count %d: | Topic(%s) | Message(%s) \n", msgCount, msg.Topic, string(msg.Value))
				sendmail(msg.Value)
			case <-sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	fmt.Println("Processed", msgCount, "messages")

	if err := worker.Close(); err != nil {
		panic(err)
	}

}

// function to connect to kafka cluster
func connectConsumer(brokersUrl []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Create new consumer
	conn, err := sarama.NewConsumer(brokersUrl, config)
	if err != nil {

		return nil, err
	}

	return conn, nil
}

func sendmail(msg []byte) {
	fmt.Println("Sending email to users", string(msg))
	resendAPIkey := "re_QAUNKZGL_856ZGuMG4aMpDy8r31GSLN4v"
	kafkaMsg := kafkaMsg{}

	err := json.Unmarshal(msg, &kafkaMsg)

	if err != nil {
		log.Println("Error unmarshalling json", err.Error())
	}

	client := resend.NewClient(resendAPIkey)

	params := &resend.SendEmailRequest{
		From:    "senuram88@gmail.com",
		To:      []string{"senuram86@gmail.com"},
		Subject: kafkaMsg.Message,
		Text:    kafkaMsg.Description,
	}
	_, err = client.Emails.Send(params)
	if err != nil {
		log.Println("failed to send email request to resend api", err.Error())
		return
	}
	fmt.Println("email sent to users")
}
