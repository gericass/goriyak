package public

import (
	"time"
	"encoding/json"
	"github.com/gericass/goriyak/model"
	"net/http"
	"io/ioutil"
)

// Admin : struct of administration node
type Admin struct {
	ID       string    `json:"id"`
	IP       string    `json:"ip"`
	Status   string    `json:"status"`
	JoinedAt *time.Time `json:"joined_at"`
}

// GetAdmin : method for get administration node by key(ID)
func GetAdmin(key string) (*Admin, error) {
	url := baseURL + "/buckets/admin/keys/" + key
	res, err := model.GetRequest(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, model.HTTPError(res)
	}
	jsonBytes, _ := ioutil.ReadAll(res.Body)
	admin := new(Admin)
	err = json.Unmarshal(jsonBytes, admin)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

// PutAdmin : method for put new administration node to riak
func (p *Admin) PutAdmin() error {
	transaction, err := json.Marshal(p)
	if err != nil {
		return err
	}
	url := baseURL + "/buckets/admin/keys/" + p.ID
	res, err := model.PutRequest(url, string(transaction))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusNoContent {
		return model.HTTPError(res)
	}
	return nil
}
