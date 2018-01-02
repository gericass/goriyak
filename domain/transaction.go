package domain

import (
	pb "github.com/gericass/goriyak/proto"
	"github.com/gericass/goriyak/model/local"
	"database/sql"
)

func ClientTransactionRequestController(r *pb.TransactionRequest, db *sql.DB) error {
	exists, err := local.GetTransactionExists(r.Name, db)
	if err != nil {
		return err
	}
	if exists {
		// TODO implements Exist Flow
	} else {
		// TODO implements Not Exist Flow
	}
	return err
}
