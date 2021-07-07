package kafka

import "github.com/confluentinc/confluent-kafka-go/kafka"

type KafkaProducer struct {
	Producer *kafka.Producer
}

func NewKafkaProducer() *KafkaProducer {
	return &KafkaProducer{}
}

func (k *KafkaProducer) SetupProducer(bootstrapServer string) {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": bootstrapServer,
	}
	k.Producer, _ = kafka.NewProducer(configMap)
}

func (k *KafkaProducer) Publish(msg, topic string) error {
	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: []byte(msg),
	}

	err := k.Producer.Produce(message, nil)
	if err != nil {
		return err
	}

	return nil
}
