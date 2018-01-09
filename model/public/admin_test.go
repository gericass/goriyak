package public

import (
	"testing"
	"time"
)

func TestGetAdmin(t *testing.T) {

	res, err := GetAdmin("test")
	if err != nil {
		t.Error("error: ", err)
	}
	expected := "test5"
	actual := res.ID
	if expected != actual {
		t.Errorf("expected %s, actual %s", expected, actual)
	}
}

func TestPutAdmin(t *testing.T) {
	tr := &Admin{ID: "test",
		IP:       "192.34.12.3",
		Status:   "active",
		JoinedAt: time.Now(),
	}
	err := tr.PutAdmin()
	if err != nil {
		t.Error("error: ", err)
	}
}
