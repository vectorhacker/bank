package events_test

import (
	"reflect"
	"testing"

	"github.com/satori/go.uuid"
	"github.com/vectorhacker/bank/core/events"
)

type MessageSent struct {
	events.Model
	Message  string `json:"message"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
}

func TestJSONSerializer(t *testing.T) {
	t.Run("it must return same event", func(t *testing.T) {
		expected := &MessageSent{
			Model: events.Model{
				EventAggregateID: uuid.Must(uuid.NewV4()),
			},
		}

		serializer := events.NewJSONSerializer(&MessageSent{})

		record, err := serializer.Serialize(expected)
		if err != nil {
			t.Fatal(err)
		}

		got, err := serializer.Deserialize(record)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(expected, got) {
			t.Fatal("Expected", expected, "got", got)
		}
	})
}
