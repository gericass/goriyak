package public

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// PublicBlock : bind the json of block for riak
type PublicBlock struct {
	ID             string    `json:"id"`
	TransactionIds []string  `json:"transaction_ids"`
	StartedAt      time.Time `json:"started_at"`
	FinishedAt     time.Time `json:"finished_at"`
	Sign           []string  `json:"sign"`
	Nonce          string    `json:"nonce"`
	CreatedAt      time.Time `json:"created_at"`
	Difficulty     string    `json:"difficulty"`
	Hash           string    `json:"hash"`
	PreviousHash   string    `json:"previous_hash"`
}

// PutBlock : method for put new block to riak
func (p *PublicBlock) PutBlock() error {
	transaction, err := json.Marshal(p)
	if err != nil {
		return err
	}
	url := baseURL + "/buckets/block/keys/" + p.ID
	res, err := PutRequest(url, string(transaction))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusNoContent {
		return HTTPError(res)
	}
	return nil
}

// GetBlock : method for get block by key(ID)
func GetBlock(key string) (*PublicBlock, error) {
	url := baseURL + "/buckets/block/keys/" + key
	res, err := GetRequest(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, HTTPError(res)
	}
	jsonBytes, _ := ioutil.ReadAll(res.Body)
	block := new(PublicBlock)
	err = json.Unmarshal(jsonBytes, block)
	if err != nil {
		return nil, err
	}
	return block, nil
}
