package events

import (
	"encoding/json"

	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
)

// Dispatcher is responsible for sending messages to the event log
type Dispatcher interface {
	Dispatch(...Event) error
}

// KafkaDispatcher dispatches events to the kafka event log
type KafkaDispatcher struct {
	serializer Serializer
	producer   sarama.SyncProducer
	topic      string
}

// NewKafkaDispatcher creates a new dispatcher for Kafka
func NewKafkaDispatcher(brokerList []string, topic string, serializer Serializer) (*KafkaDispatcher, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 10
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		return nil, err
	}

	return &KafkaDispatcher{
		topic:      topic,
		producer:   producer,
		serializer: serializer,
	}, nil
}

// Dispatch dispatches messages to the event log
func (d *KafkaDispatcher) Dispatch(events ...Event) error {

	messages := make([]*sarama.ProducerMessage, len(events))
	for i, event := range events {
		record, err := d.serializer.Serialize(event)
		if err != nil {
			return errors.Wrap(err, "unable to serialize event")
		}

		data, err := json.Marshal(record)
		if err != nil {
			return errors.Wrap(err, "unable to marshal record")
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
func (d *KafkaDispatcher) Close() error {
	return d.producer.Close()
}
