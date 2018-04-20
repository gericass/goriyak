package public

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
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

// GetNode : method for get node by key(ID)
func GetNode(key string) (*PublicNode, error) {
	url := baseURL + "/buckets/node/keys/" + key
	res, err := GetRequest(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, HTTPError(res)
	}
	jsonBytes, _ := ioutil.ReadAll(res.Body)
	node := new(PublicNode)
	err = json.Unmarshal(jsonBytes, node)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// PutNode : method for put new node to riak
func (p *PublicNode) PutNode() error {
	node, err := json.Marshal(p)
	if err != nil {
		return err
	}
	url := baseURL + "/buckets/node/keys/" + p.ID
	res, err := PutRequest(url, string(node))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusNoContent {
		return HTTPError(res)
	}
	return nil
}

// DeleteNode : method for delete new node to riak
func (p *PublicNode) DeleteNode() error {
	url := baseURL + "/buckets/node/keys/" + p.ID
	res, err := DeleteRequest(url)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusNoContent {
		return HTTPError(res)
	}
	return nil
}
