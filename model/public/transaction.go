package public

import (
	"time"
	"encoding/json"
	"os/exec"
)

// PublicTransaction : bind the json of transaction for riak
type PublicTransaction struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	SendNodeID    int64     `json:"send_node_id"`
	ReceiveNodeID int64     `json:"receive_node_id"`
	Amount        float64   `json:"amount"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}

// PutTransaction : method for put new transaction to riak
func (p *PublicTransaction) PutTransaction() (string, error) {
	transaction, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	url := Host + "/buckets/transaction/keys/" + p.ID + "'"
	jsonString := "'" + string(transaction) + "'"
	out, err := exec.Command(ComCurl, OptX, OptPUT, OptI, url, OptH, OptJson, OptD, jsonString).CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
