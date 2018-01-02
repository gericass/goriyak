package main

import (
	"net"
	"google.golang.org/grpc"
	pb "github.com/gericass/goriyak/proto"
	"log"
	"github.com/gericass/goriyak/application/handler"
	"database/sql"
)

const grpcPort = ":50051"

func connectDB() (*sql.DB, error) {
	cnn, err := sql.Open("mysql", "root:mysql@tcp(127.0.0.1:13306)/goriyak?parseTime=true")
	if err != nil {
		return nil, err
	}
	return cnn, nil
}

func main() {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Println("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	db, err := connectDB()
	if err != nil {
		log.Println("failed to connect Database")
	}
	server := &handler.GoriyakServer{DB: db}
	pb.RegisterAdminServer(s, server)
	pb.RegisterGoriyakServer(s, server)
	s.Serve(lis)
}
