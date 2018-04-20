package domain

import (
	"database/sql"
	"time"

	"github.com/gericass/goriyak/db/local"
	pb "github.com/gericass/goriyak/proto"
)

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
