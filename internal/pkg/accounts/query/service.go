package query

import (
	"github.com/jinzhu/gorm"
	pb "github.com/vectorhacker/bank/pb/accounts"
)

type Service struct {
	db *gorm.DB
}

func (s *Service) ListAccounts(r *pb.ListAccountsRequest, server pb.AccountsQuery_ListAccountsServer) error {
	return nil
}

func (s *Service) ListTransactions(r *pb.ListTransactionsRequest, server pb.AccountsQuery_ListTransactionsServer) error {
	return nil
}
