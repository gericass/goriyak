package main

import (
	"net"
	"google.golang.org/grpc"
	pb "github.com/gericass/goriyak/proto"
	"log"
	"github.com/gericass/goriyak/application/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gericass/goriyak/model/local"
)

const grpcPort = ":50051"

func main() {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Println("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	db, err := local.ConnectDB()
	if err != nil {
		log.Println("failed to connect DB: ", err)
	}
	defer db.Close()
	server := &handler.GoriyakServer{DB: db}
	pb.RegisterAdminServer(s, server)
	pb.RegisterGoriyakServer(s, server)
	s.Serve(lis)
}
