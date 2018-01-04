package domain

import (
	"database/sql"
	pb "github.com/gericass/goriyak/proto"
	"github.com/gericass/goriyak/model/local"
	"crypto/sha256"
	"encoding/hex"
	"github.com/golang/protobuf/ptypes"
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

func AdminTransactionRequestController(r *pb.Transaction, db *sql.DB) error {

	if r.Status == "approved" {
		if err := saveTransactionLocal(r, db); err != nil {
			return err
		}
	} else {
		// TODO implements flow that transaction is unapproved
	}
	return nil
}
