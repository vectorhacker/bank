package events

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

type Handler interface {
	On(Event) error
}

type HandlerFunc func(Event) error

func (h HandlerFunc) On(event Event) error {
	return h(event)
}

type Consumer struct {
	handlers     []Handler
	deserializer Deserializer
	brokers      []string
	topics       []string
	groupName    string
	config       *cluster.Config
}

func NewConsumer(
	brokers []string,
	topics []string,
	groupName string,
	deserializer Deserializer,
	handlers ...Handler,
) *Consumer {
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	return &Consumer{
		config:       config,
		topics:       topics,
		brokers:      brokers,
		handlers:     handlers,
		groupName:    groupName,
		deserializer: deserializer,
	}
}

// Start starts the event consumer
func (c *Consumer) Start(ctx context.Context) error {
	log.Println("starting consumer", c.groupName)
	consumer, err := cluster.NewConsumer(c.brokers, c.groupName, c.topics, c.config)
	if err != nil {
		return err
	}
	defer consumer.Close()

	for {
		select {
		case msg := <-consumer.Messages():
			record := Record{}

			err := json.Unmarshal(msg.Value, &record)
			if err != nil {
				return err
			}

			// deserialize record into an event
			event, err := c.deserializer.Deserialize(record)
			if err != nil {
				switch err {
				case ErrCannotDeserializeUnkown:
					// skip on unkonwn event type
					continue
				default:
					return nil
				}
			}

			// send event to event handlers
			for _, h := range c.handlers {
				if err := h.On(event); err != nil {
					return err
				}
			}

			// only mark offset if successfully processed message
			consumer.MarkOffset(msg, "")
		case <-ctx.Done():
			return nil
		}
	}
}
