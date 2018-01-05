package handler

import (
	"golang.org/x/net/context"
	pb "github.com/gericass/goriyak/proto"
	"github.com/gericass/goriyak/domain"
)

// PostBlock : to post block for confirm
func (s *GoriyakServer) PostBlock(c context.Context, r *pb.MiningResult) (*pb.Status, error) {

	return &pb.Status{Message: "Block received"}, nil
}

// PostTransaction : to post transaction for confirm transaction
func (s *GoriyakServer) PostTransactionFromServer(c context.Context, r *pb.Transaction) (*pb.Status, error) {
	if err := domain.AdminTransactionRequestController(r, s.DB); err != nil {
		return nil, err
	}
	return &pb.Status{Message: "Transaction received"}, nil
}
