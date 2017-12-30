package public

import (
	"testing"
)

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
