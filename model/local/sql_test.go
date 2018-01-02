package local

import (
	"testing"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"time"
)

func TestPutTransaction(t *testing.T) {
	cnn, err := sql.Open("mysql", "root:mysql@tcp(host:3306)/goriyak")
	if err != nil {
		t.Error("connection error")
	}
	defer cnn.Close()
	tx, err := cnn.Begin()
	tr := &LocalTransaction{Name: "testTransaction", SendNodeID: "node1", ReceiveNodeID: "node2", Amount: 16.00, CreatedAt: time.Now()}
	err = tr.PutTransaction(tx)
	if err != nil {
		t.Error(err)
	}
}
