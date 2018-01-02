package main

import (
	"net"
	"google.golang.org/grpc"
	pb "github.com/gericass/goriyak/proto"
	"log"
	"github.com/gericass/goriyak/application/handler"
)

const grpcPort = ":50051"

func main() {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Println("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAdminServer(s, handler.NewAdminServer())
	pb.RegisterGoriyakServer(s, handler.NewGoriyakServer())
	s.Serve(lis)
}
