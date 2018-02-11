package handler

import (
	"github.com/gericass/goriyak/domain"
	pb "github.com/gericass/goriyak/proto"
	"golang.org/x/net/context"
)

// PostBlock : to post block for confirm
func (s *GoriyakServer) PostBlock(c context.Context, r *pb.MiningResult) (*pb.Status, error) {
	status, err := domain.MiningController(r, s.DB)
	if err != nil {
		return status, err
	}
	return status, nil
}

// PostTransaction : to post transaction for confirm transaction
func (s *GoriyakServer) PostTransactionFromServer(c context.Context, r *pb.Transaction) (*pb.Status, error) {
	if err := domain.AdminTransactionRequestController(r, s.DB); err != nil {
		return nil, err
	}
	return &pb.Status{Message: "Transaction received"}, nil
}
