package domain

import (
	"fmt"
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	currentTime := time.Now().UTC()
	res, _ := parseTime(currentTime)
	fmt.Printf("start: %v, end: %v \n", res.start, res.end)
}
