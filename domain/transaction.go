package domain

import (
	pb "github.com/gericass/goriyak/proto"
	"github.com/gericass/goriyak/model/local"
)

func CheckTransactionExists(r *pb.TransactionRequest) bool {
	local.GetTransactionsByName(r.Name)
	return false
}
