package public

import (
	"time"
	"encoding/json"
	"github.com/gericass/goriyak/model"
	"fmt"
	"net/http"
	"bytes"
	"errors"
)

// PublicNode : bind the json of node for riak
type PublicNode struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Password       string    `json:"password"`
	Salt           string    `json:"salt"`
	JoinedAt       time.Time `json:"joined_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	ParentServerID string    `json:"parent_server_id"`
}

// TODO implement GetNode

// PutNode : method for put new node to riak
func (p *PublicNode) PutNode() error {
	node, err := json.Marshal(p)
	if err != nil {
		return err
	}
	url := baseURL + "/buckets/node/keys/" + p.ID
	res, err := model.PutRequest(url, string(node))
	if err != nil {
		return err
	}
	fmt.Println(res.StatusCode)
	if res.StatusCode != http.StatusNoContent {
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		return errors.New(buf.String())
	}
	return nil
}

// DeleteNode : method for delete new node to riak
func (p *PublicNode) DeleteNode() error {
	url := baseURL + "/buckets/node/keys/" + p.ID
	res, err := model.DeleteRequest(url)
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
