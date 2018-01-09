package domain

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"github.com/gericass/goriyak/model/local"
	"github.com/gericass/goriyak/model/public"
	pb "github.com/gericass/goriyak/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"time"
)

func generateUniqueKeyClient(r *pb.TransactionRequest) string {
	key := r.Name + r.SendNodeId + r.ReceiveNodeId
	hash := sha256.Sum256([]byte(key))
	return hex.EncodeToString(hash[:])
}

func saveTransactionToRiak(r *pb.TransactionRequest, currentTime time.Time) error {
	tr := &public.PublicTransaction{
		ID:            generateUniqueKeyClient(r),
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

func transferTransaction(r *pb.TransactionRequest) error {
	receiveNode, err := public.GetNode(r.ReceiveNodeId)
	if err != nil {
		return err
	}
	admin, err := public.GetAdmin(receiveNode.ParentServerID)
	if err != nil {
		return err
	}
	conn, err := grpc.Dial(admin.IP+":50051", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	c := pb.NewAdminClient(conn)

	tr := &pb.Transaction{
		Name:          r.Name,
		SendNodeId:    r.SendNodeId,
		ReceiveNodeId: r.ReceiveNodeId,
		Amount:        r.Amount,
		CreatedAt:     r.CreatedAt,
	}

	if _, err := c.PostTransactionFromServer(context.Background(), tr); err != nil {
		return err
	}
	return nil
}

// ClientTransactionRequestController : handle the transaction sent from client
func ClientTransactionRequestController(r *pb.TransactionRequest, db *sql.DB) error {
	exists, err := local.GetTransactionExists(r.Name, db)
	if err != nil {
		return err
	}
	currentTime := time.Now().UTC()

	if exists {
		if err := saveTransactionToRiak(r, currentTime); err != nil {
			return err
		}
		if err := local.UpdateTransactionStatus(r.Name, currentTime, db); err != nil {
			return err
		}
		if err := MulticastTransactionClient(r, currentTime); err != nil {
			return err
		}
	} else {
		if err := transferTransaction(r); err != nil {
			return err
		}
	}
	return err
}
