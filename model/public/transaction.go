package public

import (
	"time"
	"encoding/json"
	"net/http"
	"errors"
	"bytes"
)

// PublicTransaction : bind the json of transaction for riak
type PublicTransaction struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	SendNodeID    string    `json:"send_node_id"`
	ReceiveNodeID string    `json:"receive_node_id"`
	Amount        float64   `json:"amount"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}

// PutTransaction : method for put new transaction to riak
func (p *PublicTransaction) PutTransaction() error {
	transaction, err := json.Marshal(p)
	if err != nil {
		return err
	}
	url := baseURL + "/buckets/transaction/keys/" + p.ID
	res, err := PutRequest(url, string(transaction))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusNoContent {
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		return errors.New(buf.String())
	}
	return nil
}
