package events

import (
	"encoding/json"

	"github.com/Shopify/sarama"
)

// Dispatcher is responsible for sending messages to the event log
type Dispatcher struct {
	serializer Serializer
	producer   sarama.SyncProducer
	topic      string
}

// NewDispatcher creates a new dispatcher
func NewDispatcher(brokerList []string, topic string, serializer Serializer) (*Dispatcher, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		return nil, err
	}

	return &Dispatcher{
		topic:      topic,
		producer:   producer,
		serializer: serializer,
	}, nil
}

// Dispatch dispatches messages to the event log
func (d Dispatcher) Dispatch(events ...Event) error {

	messages := make([]*sarama.ProducerMessage, len(events))
	for i, event := range events {
		record, err := d.serializer.Serialize(event)
		if err != nil {
			return err
		}

		data, err := json.Marshal(record)
		if err != nil {
			return err
		}

		message := &sarama.ProducerMessage{
			Topic: d.topic,
			Key:   sarama.ByteEncoder(event.AggregateID().Bytes()),
			Value: sarama.ByteEncoder(data),
		}

		messages[i] = message
	}

	return d.producer.SendMessages(messages)
}

// Close closes the connection to the event log
func (d Dispatcher) Close() error {
	return d.producer.Close()
}
