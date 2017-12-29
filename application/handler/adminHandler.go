package handler

import (
	"golang.org/x/net/context"
	pb "github.com/gericass/goriyak/proto"
)

// AdminServer : empty struct for implements proto.AdminServer
type AdminServer struct{}

// NewAdminServer : for register AdminServer
func NewAdminServer() *AdminServer {
	return &AdminServer{}
}

// PostBlock : to post block for confirm
func (s *AdminServer) PostBlock(c context.Context, r *pb.MiningResult) (*pb.Status, error) {

	return &pb.Status{Message: "Block received"}, nil
}

// PostTransaction : to post transaction for confirm transaction
func (s *AdminServer) PostTransaction(c context.Context, r *pb.Transaction) (*pb.Status, error) {

	return &pb.Status{Message: "Transaction received"}, nil
}
