package transfers

import (
	"context"
	"time"

	"github.com/satori/go.uuid"

	"github.com/vectorhacker/bank/internal/pkg/events"
	td "github.com/vectorhacker/bank/internal/pkg/events/transfers"
	pb "github.com/vectorhacker/bank/pb/transfers"
)

type Service struct {
	dispatcher events.Dispatcher
}

func NewService(dispatcher events.Dispatcher) *Service {
	return &Service{
		dispatcher: dispatcher,
	}
}

// BeginTransfer begins a transfer
func (s *Service) BeginTransfer(
	ctx context.Context,
	r *pb.BeginTransferRequest,
) (*pb.BeginTransferResponse, error) {

	id := uuid.Must(uuid.NewV4())

	from, err := uuid.FromString(r.FromAccountID)
	to, err := uuid.FromString(r.ToAccountID)
	if err != nil {
		return nil, err
	}

	err = s.dispatcher.Dispatch(&td.TransferBegun{
		Model: events.Model{
			EventAggregateID: id,
			EventAt:          time.Now(),
			EventID:          uuid.Must(uuid.NewV4()),
		},
		Amount:      r.Amount,
		Description: r.Description,
		FromAccount: from,
		ToAccount:   to,
	})
	if err != nil {
		return nil, err
	}

	return &pb.BeginTransferResponse{
		ID: id.String(),
	}, nil
}
