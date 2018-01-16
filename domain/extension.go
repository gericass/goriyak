package domain

import (
	"encoding/base64"
	"time"
)

type timeSet struct {
	start time.Time
	end   time.Time
}

func (t *timeSet) generateBlockID() string {
	seed := t.start.String() + " : " + t.end.String()
	str := base64.StdEncoding.EncodeToString([]byte(seed))
	return str
}
