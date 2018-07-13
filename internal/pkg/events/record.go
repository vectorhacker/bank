package events

import (
	"encoding/json"
	"errors"
	"reflect"
)

// Errors
var (
	ErrCannotDeserializeUnkown = errors.New("cannot deserialize unknown type")
)

// Record gets saved to event store (kafka)
type Record struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

// Serializer serializes Event into a record
type Serializer interface {
	Serialize(Event) (Record, error)
}

// Deserializer deserializes a record into an event
type Deserializer interface {
	Deserialize(Record) (Event, error)
}

// JSONSerializer implements the Serializer and Deserializer interfaces
type JSONSerializer struct {
	types map[string]reflect.Type
}

// NewJSONSerializer creates a JSONSerializer
func NewJSONSerializer(events ...Event) *JSONSerializer {
	types := make(map[string]reflect.Type, len(events))

	for _, event := range events {
		t := reflect.TypeOf(event)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}

		types[t.Name()] = t
	}

	return &JSONSerializer{
		types: types,
	}
}

// Deserialize implements the Deserializer interface
func (s JSONSerializer) Deserialize(r Record) (Event, error) {
	v, ok := s.types[r.Type]
	if !ok {
		return nil, ErrCannotDeserializeUnkown
	}

	ev := reflect.New(v).Interface().(Event)
	if err := json.Unmarshal(r.Payload, ev); err != nil {
		return nil, err
	}

	return ev, nil
}

// Serialize implements the Serializer interface
// additionally, adds event type to type map if it wasn't added before
func (s *JSONSerializer) Serialize(ev Event) (Record, error) {
	data, err := json.Marshal(ev)
	if err != nil {
		return Record{}, err
	}

	t := reflect.TypeOf(ev)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	eventType := t.Name()

	// add event type to type map if not exists
	if _, ok := s.types[eventType]; !ok {
		s.types[eventType] = t
	}

	return Record{
		Payload: data,
		Type:    eventType,
	}, nil
}
