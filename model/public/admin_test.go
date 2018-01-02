package public

import (
	"testing"
	"time"
)

func TestPutAdmin(t *testing.T) {
	tr := &Admin{ID: "test",
		IP: "192.34.12.3",
		Status: "active",
		JoinedAt: time.Now(),
	}
	err := tr.PutAdmin()
	if err != nil {
		t.Error("error: ", err)
	}
}
