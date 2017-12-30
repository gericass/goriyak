package public

import (
	"testing"
	"time"
	"strings"
)

func TestPutTransaction(t *testing.T) {
	tr := &PublicTransaction{ID: "test",
		Name: "testNode",
		SendNodeID: "send",
		ReceiveNodeID: "receive",
		Amount: 0.00,
		Status: "approved",
		CreatedAt: time.Now(),
	}
	actual, _ := tr.PutTransaction()
	expected := "204"
	if !strings.Contains(actual, expected) {
		t.Error("PutTransaction(): expected %s, actual %s", expected, actual)
	}
}
