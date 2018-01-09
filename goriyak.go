package main

import (
	"database/sql"
	"github.com/gericass/goriyak/application/handler"
	"github.com/gericass/goriyak/model/local"
	pb "github.com/gericass/goriyak/proto"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

const grpcPort = ":50051"

func gracefulShutdown(db *sql.DB, signalChan chan os.Signal) {
	for {
		s := <-signalChan
		switch s {
		case syscall.SIGINT:
			db.Close()
			if err := exec.Command("riak-admin", "cluster", "leave").Run(); err != nil {
				log.Printf("riak error: %v\n", err)
			}
			log.Println("leaved")
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Printf("failed to listen: %v\n", err)
	}
	s := grpc.NewServer()
	db, err := local.ConnectDB()
	if err != nil {
		log.Printf("failed to connect DB: %v ", err)
	}
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)
	go gracefulShutdown(db, signalChan)
	defer db.Close()
	server := &handler.GoriyakServer{DB: db}
	pb.RegisterAdminServer(s, server)
	pb.RegisterGoriyakServer(s, server)
	s.Serve(lis)
}
