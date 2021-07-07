package protocols

type KafkaProducer interface {
	Publish(msg, topic string) error
}
