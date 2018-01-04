package domain

import (
	pb "github.com/gericass/goriyak/proto"
	"github.com/gericass/goriyak/model/local"
	"database/sql"
	"github.com/gericass/goriyak/model/public"
	"crypto/sha256"
	"encoding/hex"
	"time"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/golang/protobuf/ptypes"
)

func generateUniqueKey(r *pb.TransactionRequest) string {
	key := r.Name + r.SendNodeId + r.ReceiveNodeId
	hash := sha256.Sum256([]byte(key))
	return hex.EncodeToString(hash[:])
}

func saveTransactionToRiak(r *pb.TransactionRequest, currentTime time.Time) error {
	tr := &public.PublicTransaction{
		ID:            generateUniqueKey(r),
		Name:          r.Name,
		SendNodeID:    r.SendNodeId,
		ReceiveNodeID: r.ReceiveNodeId,
		Amount:        r.Amount,
		Status:        "approved",
		CreatedAt:     currentTime,
	}
	if err := tr.PutTransaction(); err != nil {
		return err
	}
	return nil
}

func getAdminIPs() ([]*public.Admin, error) {
	keys, err := public.GetAdminKey()
	if err != nil {
		return []*public.Admin{}, nil
	}
	admins := make([]*public.Admin, 0)
	for _, v := range keys.Keys {
		admin, err := public.GetAdmin(v)
		if err != nil {
			return []*public.Admin{}, err
		}
		admins = append(admins, admin)
	}
	return admins, nil
}

func broadcastTransaction(ip string, tr *pb.Transaction) error {
	conn, err := grpc.Dial(ip+":50051", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	c := pb.NewAdminClient(conn)

	if _, err := c.PostTransactionFromServer(context.Background(), tr); err != nil {
		return err
	}
	return nil
}

func multicastTransaction(r *pb.TransactionRequest, currentTime time.Time) error {
	admins, err := getAdminIPs()
	if err != nil {
		return err
	}
	timeProto, err := ptypes.TimestampProto(currentTime)
	if err != nil {
		return err
	}
	tr := &pb.Transaction{
		Name:          r.Name,
		SendNodeId:    r.SendNodeId,
		ReceiveNodeId: r.ReceiveNodeId,
		Amount:        r.Amount,
		CreatedAt:     timeProto,
	}
	for _, v := range admins {
		err := broadcastTransaction(v.IP, tr)
		if err != nil {
			return err
		}
	}
	return nil

}

func ClientTransactionRequestController(r *pb.TransactionRequest, db *sql.DB) error {
	exists, err := local.GetTransactionExists(r.Name, db)
	if err != nil {
		return err
	}
	currentTime := time.Now().UTC()
	if exists {
		saveTransactionToRiak(r, currentTime)
		local.UpdateTransactionStatus(r.Name, currentTime, db)
		multicastTransaction(r, currentTime)
	} else {
		// TODO implements Not Exist Flow
	}
	return err
}
