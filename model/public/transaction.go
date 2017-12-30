package public

import (
	"time"
	"github.com/basho/taste-of-riak/go/util"
	"github.com/basho/riak-go-client"
	"sync"
	"encoding/json"
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
	o := &riak.NewClientOptions{
		RemoteAddresses: []string{util.GetRiakAddress()},
	}

	var c *riak.Client
	c, err := riak.NewClient(o)
	if err != nil {
		util.ErrExit(err)
	}

	defer func() {
		if err := c.Stop(); err != nil {
			util.ErrExit(err)
		}
	}()

	// NB: ensure that members are exported (i.e. capitalized)
	var jsonbytes []byte
	jsonbytes, err = json.Marshal(p)
	if err != nil {
		util.ErrExit(err)
	}

	objs := []*riak.Object{
		{
			Bucket:      "transaction",
			Key:         p.ID,
			ContentType: "application/json",
			Value:       jsonbytes,
		},
	}

	var cmd riak.Command
	wg := &sync.WaitGroup{}

	for _, o := range objs {
		cmd, err = riak.NewStoreValueCommandBuilder().
			WithContent(o).
			Build()
		if err != nil {
			return err
		}
		a := &riak.Async{
			Command: cmd,
			Wait:    wg,
		}
		if err := c.ExecuteAsync(a); err != nil {
			util.ErrLog.Println(err)
		}
	}

	wg.Wait()
	return nil
}
