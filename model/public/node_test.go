package public

import (
	"testing"
)

func TestGetNode(t *testing.T) {

	res, err := GetNode("test5")
	if err != nil {
		t.Error("error: ", err)
	}
	expected := "test5"
	actual := res.ID
	if expected != actual {
		t.Errorf("expected %s, actual %s", expected, actual)
	}
}

func TestPutNode(t *testing.T) {
	tr := &PublicNode{ID: "test5"}
	err := tr.PutNode()
	if err != nil {
		t.Error("error: ", err)
	}
}

func TestDeleteNode(t *testing.T) {
	tr := &PublicNode{ID: "test5"}
	err := tr.DeleteNode()
	if err != nil {
		t.Error("error: ", err)
	}
}
