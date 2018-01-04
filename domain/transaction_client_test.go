package domain

import (
	"testing"
	pb "github.com/gericass/goriyak/proto"
	"time"
	"github.com/golang/protobuf/ptypes"
	"database/sql"
)

func TestClientTransactionRequestController(t *testing.T) {
	tm, err := ptypes.TimestampProto(time.Now().UTC())
	if err != nil {
		t.Errorf("error: ", err)
	}
	cnn, err := sql.Open("mysql", "root:mysql@tcp(127.0.0.1:13306)/goriyak?parseTime=true")
	if err != nil {
		t.Error("connection error")
	}
	defer cnn.Close()

	tr := &pb.TransactionRequest{
		Name:          "tr1",
		SendNodeId:    "node1",
		ReceiveNodeId: "node2",
		Amount:        12.23,
		CreatedAt:     tm,
	}

	err = ClientTransactionRequestController(tr, cnn)
	if err != nil {
		t.Error("error: ", err)
	}

}
