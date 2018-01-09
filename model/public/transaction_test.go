package public

import (
	"testing"
	"time"
)

func TestPutTransaction(t *testing.T) {
	tr := &PublicTransaction{ID: "test5",
		Name:          "testNode",
		SendNodeID:    "send",
		ReceiveNodeID: "receive",
		Amount:        0.00,
		Status:        "approved",
		CreatedAt:     time.Now(),
	}
	err := tr.PutTransaction()
	if err != nil {
		t.Error("error: ", err)
	}
}
