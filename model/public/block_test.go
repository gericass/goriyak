package public

import (
	"testing"
	"time"
)

func TestPutBlock(t *testing.T) {
	tr := &PublicBlock{ID: "test",
		TransactionIds: []string{"test1", "test2", "test3"},
		StartedAt: time.Now(),
		FinishedAt: time.Now(),
		Sign: []string{"test", "test2", "test3"},
		PreviousHash: "test",
		Nonce: "test",
		CreatedAt: time.Now(),
		Difficulty: "test",
	}
	err := tr.PutBlock()
	if err != nil {
		t.Error("error: ", err)
	}
}
