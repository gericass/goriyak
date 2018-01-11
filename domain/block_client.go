package domain

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"github.com/gericass/goriyak/model/local"
	"github.com/gericass/goriyak/model/public"
	pb "github.com/gericass/goriyak/proto"
	"github.com/golang/protobuf/ptypes"
	"strconv"
	"time"
	"encoding/base64"
)

type timeSet struct {
	start time.Time
	end   time.Time
}

func generateTransactionID(r *local.LocalTransaction) string {
	key := r.Name + r.SendNodeID + r.ReceiveNodeID
	hash := sha256.Sum256([]byte(key))
	return hex.EncodeToString(hash[:])
}

func (t *timeSet) generateBlockID() string {
	seed := t.start.String() + " : " + t.end.String()
	str := base64.StdEncoding.EncodeToString([]byte(seed))
	return str
}

func parseTime(currentTime time.Time) (*timeSet, error) {
	t := currentTime.Minute() % 5
	var start time.Time
	var end time.Time
	minusFive, err := time.ParseDuration("-5m")
	if err != nil {
		return nil, err
	}
	if t != 0 {
		d, err := time.ParseDuration("-" + strconv.Itoa(t) + "m")
		if err != nil {
			return nil, err
		}
		end = currentTime.Add(d).Truncate(time.Minute)
		start = end.Add(minusFive)

	} else {
		end = currentTime.Truncate(time.Minute)
		start = end.Add(minusFive)
	}
	return &timeSet{start: start, end: end}, nil
}

func (t *timeSet) getPreviousHash() (string, error) {
	minusFive, err := time.ParseDuration("-5m")
	if err != nil {
		return "", err
	}
	previousEnd := t.end.Add(minusFive)
	previousStart := t.start.Add(minusFive)
	previousTimeSet := &timeSet{
		start: previousStart,
		end:   previousEnd,
	}
	block, err := public.GetBlock(previousTimeSet.generateBlockID())
	if err != nil {
		return "", err
	}
	return block.Hash, nil
}

func (t *timeSet) generateBlock(trs []*local.LocalTransaction) (*pb.Block, error) {
	blockTransactions := make([]*pb.Block_Transaction, 0)
	startProto, err := ptypes.TimestampProto(t.start)
	if err != nil {
		return nil, err
	}
	endProto, err := ptypes.TimestampProto(t.end)
	if err != nil {
		return nil, err
	}
	for _, v := range trs {
		timeProto, err := ptypes.TimestampProto(v.UpdatedAt)
		if err != nil {
			return nil, err
		}
		bltr := &pb.Block_Transaction{
			Id:            generateTransactionID(v),
			SendNodeId:    v.SendNodeID,
			ReceiveNodeId: v.ReceiveNodeID,
			Amount:        v.Amount,
			CreatedAt:     timeProto,
		}
		blockTransactions = append(blockTransactions, bltr)
	}
	previousHash, err := t.getPreviousHash()
	if err != nil {
		return nil, err
	}
	block := &pb.Block{
		Id:           t.generateBlockID(),
		Transactions: blockTransactions,
		StartedAt:    startProto,
		FinishedAt:   endProto,
		PreviousHash: previousHash,
		Difficulty:   "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF", // FIXME set hash
	}
	return block, nil
}

// ClientBlockController : generate block
func ClientBlockController(db *sql.DB) (*pb.Block, error) {
	currentTime := time.Now().UTC()
	ts, err := parseTime(currentTime)
	if err != nil {
		return nil, err
	}
	trs, err := local.GetTransactionsByTime(ts.start, ts.end, db)
	if err != nil {
		return nil, err
	}
	block, err := ts.generateBlock(trs)
	if err != nil {
		return nil, err
	}
	return block, nil
}
