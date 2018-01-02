package public

import (
	"time"
	"encoding/json"
	"github.com/gericass/goriyak/model"
	"net/http"
)

// Admin : struct of administration node
type Admin struct {
	ID       string    `json:"id"`
	IP       string    `json:"ip"`
	Status   string    `json:"status"`
	JoinedAt time.Time `json:"joined_at"`
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
