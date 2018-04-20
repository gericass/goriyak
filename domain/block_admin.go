package domain

import (
	"database/sql"
	"encoding/base64"
	"strings"
	"time"

	"fmt"

	"crypto/sha256"

	"github.com/golang/protobuf/ptypes"
	"github.com/gericass/goriyak/db/local"
	"github.com/gericass/goriyak/db/public"
	pb "github.com/gericass/goriyak/proto"
	"github.com/gericass/goriyak/setting"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// post MiningResult to other admin node
func broadcastMiningResult(r *pb.MiningResult) error {

	admins, err := public.GetAdminKey()
	if err != nil {
		return err
	}
	var sendTo string
	for _, vk := range admins.Keys {
		flag := true
		for _, vc := range r.Check {
			if vk == vc {
				flag = false
			}
		}
		if flag {
			sendTo = vk
		}
	}

	admin, err := public.GetAdmin(sendTo)
	conn, err := grpc.Dial(admin.IP+":50051", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	c := pb.NewAdminClient(conn)

	if _, err := c.PostBlock(context.Background(), r); err != nil {
		return err
	}
	return nil
}

func checkActiveMiningResult(r *pb.MiningResult) (bool, error) {
	admins, err := public.GetAdminKey()
	if err != nil {
		return false, err
	}
	lack := (len(admins.Keys) * 2 / 3) - len(r.Sign)
	remaining := len(admins.Keys) - len(r.Check)
	if lack > remaining {
		return false, nil
	}

	return true, nil
}

func stringToTimeSet(timeString []string) (*timeSet, error) {
	format := "2018-01-18 02:03:46.864807895 +0000 UTC"
	start, err := time.Parse(timeString[0], format)
	if err != nil {
		return nil, err
	}
	end, err := time.Parse(timeString[1], format)
	if err != nil {
		return nil, err
	}
	return &timeSet{start: start, end: end}, nil
}

func generateBlockByMiningResult(r *pb.MiningResult, db *sql.DB) (*pb.Block, error) {
	timeByte, err := base64.StdEncoding.DecodeString(r.BlockId)
	if err != nil {
		return nil, err
	}
	t := strings.Split(string(timeByte), " : ")
	ts, err := stringToTimeSet(t)
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

func transactionToStaring(tr []*pb.Block_Transaction) string {
	var str string
	for _, v := range tr {
		str += v.Id + v.SendNodeId + v.ReceiveNodeId + fmt.Sprint(v.Amount) + ptypes.TimestampString(v.CreatedAt)
	}
	return str
}

func confirmHashing(b *pb.Block) bool {
	seed := b.Id + transactionToStaring(b.Transactions) + ptypes.TimestampString(b.StartedAt) + ptypes.TimestampString(b.FinishedAt) + b.PreviousHash + b.Nonce + ptypes.TimestampString(b.CreatedAt)
	hash := sha256.Sum256([]byte(seed))
	s := string(hash[:32])
	if s < "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFF" { // FIXME difficulty
		return true
	}
	return false
}

func updateMiningResult(r *pb.MiningResult, db *sql.DB) (*pb.MiningResult, error) {
	block, err := generateBlockByMiningResult(r, db)
	if err != nil {
		return nil, err
	}

	hashingResult := confirmHashing(block)
	if hashingResult {
		r.Sign = append(r.Sign, setting.ServerConfig.Name)
	}
	return r, nil
}

// MiningController : Handler for MiningResult sent by other admin node
func MiningController(miningResult *pb.MiningResult, db *sql.DB) (*pb.Status, error) {
	if res, _ := public.GetBlock(miningResult.BlockId); res != nil {
		return &pb.Status{Message: "Block already exists"}, nil
	}

	ex, err := checkActiveMiningResult(miningResult)
	if err != nil {
		return &pb.Status{Message: "Server error"}, err
	}
	if ex {

		res, err := updateMiningResult(miningResult, db)
		if err != nil {
			return &pb.Status{Message: "Mining failed"}, nil
		}

		if err := broadcastMiningResult(res); err != nil {
			return &pb.Status{Message: "Server error"}, err
		}
	}

	return &pb.Status{Message: "mining result received"}, nil

}
