package domain

import (
	"fmt"
	"testing"
	"time"
	"encoding/base64"
)

func TestParseTime(t *testing.T) {
	currentTime := time.Now().UTC()
	res, _ := parseTime(currentTime)
	fmt.Printf("start: %v, end: %v \n", res.start, res.end)
}

func TestGenerateBlockID(t *testing.T) {
	ts := timeSet{start: time.Now().UTC(), end: time.Now().UTC()}
	res := ts.generateBlockID()
	actual, err := base64.StdEncoding.DecodeString(res)
	if err != nil {
		t.Errorf("non time: %v", err)
	}
	expect := ts.start.String() + " : " + ts.end.String()
	if expect != string(actual) {
		t.Errorf("expected: %s,\n actual: %s", expect, string(actual))
	}
}
