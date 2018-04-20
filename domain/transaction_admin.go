package domain

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"github.com/gericass/goriyak/db/local"
	"github.com/gericass/goriyak/db/public"
	pb "github.com/gericass/goriyak/proto"
	"github.com/golang/protobuf/ptypes"
	"time"
)

func generateUniqueKeyAdmin(r *pb.Transaction) string {
	key := r.Name + r.SendNodeId + r.ReceiveNodeId
	hash := sha256.Sum256([]byte(key))
	return hex.EncodeToString(hash[:])
}

func saveTransactionLocal(r *pb.Transaction, db *sql.DB) error {
	ca, err := ptypes.Timestamp(r.CreatedAt)
	if err != nil {
		return err
	}
	tr := &local.LocalTransaction{
		Name:          r.Name,
		SendNodeID:    r.SendNodeId,
		ReceiveNodeID: r.ReceiveNodeId,
		Amount:        r.Amount,
		Status:        r.Status,
		CreatedAt:     ca,
		UpdatedAt:     ca,
	}
	if err := tr.PutTransaction(db); err != nil {
		return err
	}
	return nil
}

func saveTransactionToRiakAdmin(r *pb.Transaction, currentTime time.Time) error {
	tr := &public.PublicTransaction{
		ID:            generateUniqueKeyAdmin(r),
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

// AdminTransactionRequestController : handle the transaction sent from admin
func AdminTransactionRequestController(r *pb.Transaction, db *sql.DB) error {
	currentTime := time.Now().UTC()

	if r.Status == "approved" {
		if err := saveTransactionLocal(r, db); err != nil {
			return err
		}
	} else {
		exists, err := local.GetTransactionExists(r.Name, db)
		if err != nil {
			return err
		}
		if exists {
			if err := saveTransactionToRiakAdmin(r, currentTime); err != nil {
				return err
			}
			if err := local.UpdateTransactionStatus(r.Name, currentTime, db); err != nil {
				return err
			}
			if err := MulticastTransactionAdmin(r, currentTime); err != nil {
				return err
			}
		} else {
			if err := saveTransactionLocal(r, db); err != nil {
				return err
			}
		}
	}
	return nil
}
