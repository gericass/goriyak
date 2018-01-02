package public

import (
	"time"
	"encoding/json"
	"github.com/gericass/goriyak/model"
	"net/http"
)

// PublicBlock : bind the json of block for riak
type PublicBlock struct {
	ID             string    `json:"id"`
	TransactionIds []string  `json:"transaction_ids"`
	StartedAt      *time.Time `json:"started_at"`
	FinishedAt     *time.Time `json:"finished_at"`
	Sign           []string  `json:"sign"`
	PreviousHash   string    `json:"previous_hash"`
	Nonce          string    `json:"nonce"`
	CreatedAt      *time.Time `json:"created_at"`
	Difficulty     string    `json:"difficulty"`
}

// PutBlock : method for put new block to riak
func (p *PublicBlock) PutBlock() error {
	transaction, err := json.Marshal(p)
	if err != nil {
		return err
	}
	url := baseURL + "/buckets/block/keys/" + p.ID
	res, err := model.PutRequest(url, string(transaction))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusNoContent {
		return model.HTTPError(res)
	}
	return nil
}
