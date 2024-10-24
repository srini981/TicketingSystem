package pkg

import (
	"fmt"
	"github.com/IBM/sarama"
)

func PushMsgToQueue(topic string, message []byte) error {

	producer := client.kafkaProducer
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(msg)

	if err != nil {
		return err
	}

	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)

	return nil
}
