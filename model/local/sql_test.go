package local

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"time"
)

func TestPutTransaction(t *testing.T) {
	cnn, err := sql.Open("mysql", "root:mysql@tcp(127.0.0.1:13306)/goriyak?parseTime=true")
	if err != nil {
		t.Error("connection error")
	}
	defer cnn.Close()
	tr := &LocalTransaction{Name: "testTransaction", SendNodeID: "node1", ReceiveNodeID: "node2", Amount: 16.00, Status: "unapproved", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	err = tr.PutTransaction(cnn)
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

	tr, err := GetTransactionsByName("testTransaction", cnn)
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

	err = DeleteTransactionByTime(time.Now(), cnn)
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
	currentTime := time.Now().UTC()

	err = UpdateTransactionStatus("testTransaction", currentTime, cnn)
	if err != nil {
		t.Error(err)
	}
}
