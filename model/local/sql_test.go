package local

import (
	"testing"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"time"
)

func TestPutTransaction(t *testing.T) {
	cnn, err := sql.Open("mysql", "root:mysql@tcp(127.0.0.1:13306)/goriyak?parseTime=true")
	if err != nil {
		t.Error("connection error")
	}
	defer cnn.Close()
	tx, err := cnn.Begin()
	tr := &LocalTransaction{Name: "testTransaction", SendNodeID: "node1", ReceiveNodeID: "node2", Amount: 16.00, Status: "unapproved", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	err = tr.PutTransaction(tx)
	if err != nil {
		t.Error(err)
	}
}

func TestGetTransactionsByName(t *testing.T) {
	cnn, err := sql.Open("mysql", "root:mysql@tcp(127.0.0.1:13306)/goriyak?parseTime=true")
	if err != nil {
		t.Error("connection error")
	}
	defer cnn.Close()
	tx, err := cnn.Begin()

	tr, err := GetTransactionsByName("testTransaction", tx)
	if err != nil {
		t.Error(err)
	}
	expected := "testTransaction"
	actual := tr[0].Name
	if expected != actual {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestDeleteTransactions(t *testing.T) {
	cnn, err := sql.Open("mysql", "root:mysql@tcp(127.0.0.1:13306)/goriyak?parseTime=true")
	if err != nil {
		t.Error("connection error")
	}
	defer cnn.Close()
	tx, err := cnn.Begin()

	err = DeleteTransactionByTime(time.Now(), tx)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateTransactions(t *testing.T) {
	cnn, err := sql.Open("mysql", "root:mysql@tcp(127.0.0.1:13306)/goriyak?parseTime=true")
	if err != nil {
		t.Error("connection error")
	}
	defer cnn.Close()
	tx, err := cnn.Begin()

	err = UpdateTransactionStatus("testTransaction", tx)
	if err != nil {
		t.Error(err)
	}
}
