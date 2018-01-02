package handler

import (
	"golang.org/x/net/context"
	pb "github.com/gericass/goriyak/proto"
	"database/sql"
)

// GoriyakServer : empty struct for implements proto.GoriyakServer
type GoriyakServer struct {
	DB *sql.DB
}

// RegisterNode : to register new node
func (s *GoriyakServer) RegisterNode(c context.Context, r *pb.Node) (*pb.Status, error) {

	return &pb.Status{Message: "Node Registered"}, nil
}

// DeleteNode : to delete node
func (s *GoriyakServer) DeleteNode(c context.Context, r *pb.Node) (*pb.Status, error) {

	return &pb.Status{Message: "Node Deleted"}, nil
}

// Login : Login endpoint
func (s *GoriyakServer) Login(ctx context.Context, r *pb.Node) (*pb.Status, error) {

	return &pb.Status{Message: "Login succeeded"}, nil
}

// PostTransaction : registering and approving new transaction
func (s *GoriyakServer) PostTransactionFromClient(ctx context.Context, r *pb.TransactionRequest) (*pb.Status, error) {
	// WIP
	return &pb.Status{Message: "Transaction transfer"}, nil
}

// GetBlock : returns Block for mining by client
func (s *GoriyakServer) GetBlock(ctx context.Context, r *pb.BlockRequest) (*pb.Block, error) {

	return &pb.Block{}, nil
}

// SuccessMining : post the result of mining by client
func (s *GoriyakServer) PostMiningResult(ctx context.Context, r *pb.MiningResult) (*pb.Status, error) {

	return &pb.Status{}, nil
}
