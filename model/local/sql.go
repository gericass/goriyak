package local

import (
	"time"
	"database/sql"
)

// LocalTransaction : struct for binding transaction of MySQL
type LocalTransaction struct {
	ID            int64
	Name          string
	SendNodeID    string
	ReceiveNodeID string
	Amount        float64
	CreatedAt     *time.Time
}

// GetTransactionsByTime : to get transactions by created_at column
func GetTransactionsByTime(start time.Time, end time.Time, tx *sql.Tx) ([]*LocalTransaction, error) {
	query := "SELECT `id`,`name`,`send_node_id`,`receive_node_id`,`amount`,`created_at` FROM `transaction` WHERE ? < `created_at` AND `created_at` < ?"
	rows, err := tx.Query(query, start, end)
	if err != nil {
		return nil, err
	}
	return scanTransactions(rows)
}

// GetTransactionsByName : to get transactions by name column
func GetTransactionsByName(name string, tx *sql.Tx) ([]*LocalTransaction, error) {
	query := "SELECT `id`,`name`,`send_node_id`,`receive_node_id`,`amount`,`created_at` FROM `transaction` WHERE `name` = ? LIMIT 2"
	rows, err := tx.Query(query, name)
	if err != nil {
		return nil, err
	}
	return scanTransactions(rows)
}

// PutTransaction : to put transactions to MySQL transaction table
func (t *LocalTransaction) PutTransaction(tx *sql.Tx) error {
	query := "INSERT INTO `transaction` (`id`, `name`, `send_node_id`, `receive_node_id`,`amount`,`created_at`) values(?, ?, ?, ?, ?, ?)"
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return dbTransaction(tx, func(tx *sql.Tx) error {
		query := "INSERT INTO `transaction` (`name`, `send_node_id`, `receive_node_id`,`amount`,`created_at`) values(?, ?, ?, ?, ?)"
		stmt, err := tx.Prepare(query)
		if err != nil {
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(t.Name, t.SendNodeID, t.ReceiveNodeID, t.Amount, t.CreatedAt)
		if err != nil {
			return err
		}
		return nil
	})

}
