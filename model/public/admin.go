package public

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// Admin : struct of administration node
type Admin struct {
	ID       string    `json:"id"`
	IP       string    `json:"ip"`
	Status   string    `json:"status"`
	JoinedAt time.Time `json:"joined_at"`
}

type AdminKeys struct {
	Keys []string `json:"keys"`
}

// GetAdmin : method for get administration node by key(ID)
func GetAdmin(key string) (*Admin, error) {
	url := baseURL + "/buckets/admin/keys/" + key
	res, err := GetRequest(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, HTTPError(res)
	}
	jsonBytes, _ := ioutil.ReadAll(res.Body)
	admin := new(Admin)
	err = json.Unmarshal(jsonBytes, admin)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

// GetAdminKey : method for get administration node's key
func GetAdminKey() (*AdminKeys, error) {
	url := baseURL + "/buckets/admin/keys?keys=true"
	res, err := GetRequest(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, HTTPError(res)
	}
	jsonBytes, _ := ioutil.ReadAll(res.Body)
	keys := new(AdminKeys)
	err = json.Unmarshal(jsonBytes, keys)
	if err != nil {
		return nil, err
	}
	return keys, nil
}

// PutAdmin : method for put new administration node to riak
func (p *Admin) PutAdmin() error {
	transaction, err := json.Marshal(p)
	if err != nil {
		return err
	}
	url := baseURL + "/buckets/admin/keys/" + p.ID
	res, err := PutRequest(url, string(transaction))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusNoContent {
		return HTTPError(res)
	}
	return nil
}
