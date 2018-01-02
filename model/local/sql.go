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
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// GetTransactionsByTime : to get transactions by created_at column
func GetTransactionsByTime(start time.Time, end time.Time, tx *sql.Tx) ([]*LocalTransaction, error) {
	query := "SELECT `id`,`name`,`send_node_id`,`receive_node_id`,`amount`,`status`,`created_at`,`updated_at` FROM `transaction` WHERE ? <= `updated_at` AND `updated_at` < ? AND `status` = 'approved'"
	rows, err := tx.Query(query, start, end)
	if err != nil {
		return nil, err
	}
	return scanTransactions(rows)
}

// GetTransactionsByName : to get transactions by name column
func GetTransactionsByName(name string, tx *sql.Tx) ([]*LocalTransaction, error) {
	query := "SELECT `id`,`name`,`send_node_id`,`receive_node_id`,`amount`,`status`,`created_at`,`updated_at` FROM `transaction` WHERE `name` = ? LIMIT 2"
	rows, err := tx.Query(query, name)
	if err != nil {
		return nil, err
	}
	return scanTransactions(rows)
}

// PutTransaction : to put transactions to MySQL transaction table
func (t *LocalTransaction) PutTransaction(tx *sql.Tx) error {
	return dbTransaction(tx, func(tx *sql.Tx) error {
		query := "INSERT INTO `transaction` (`name`, `send_node_id`, `receive_node_id`, `amount`, `status`, `created_at`, `updated_at`) values(?, ?, ?, ?, ?, ?, ?)"
		stmt, err := tx.Prepare(query)
		if err != nil {
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(t.Name, t.SendNodeID, t.ReceiveNodeID, t.Amount, t.Status, t.CreatedAt, t.UpdatedAt)
		if err != nil {
			return err
		}
		return nil
	})
}

// DeleteTransactionByTime : to delete transactions to MySQL transaction table
func DeleteTransactionByTime(time time.Time, tx *sql.Tx) error {
	return dbTransaction(tx, func(tx *sql.Tx) error {
		query := "DELETE FROM `transaction` WHERE `updated_at` <= ?"
		stmt, err := tx.Prepare(query)
		if err != nil {
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(time)
		if err != nil {
			return err
		}
		return nil
	})
}

// UpdateTransactionStatus : to update transaction's status
func UpdateTransactionStatus(name string, tx *sql.Tx) error {
	return dbTransaction(tx, func(tx *sql.Tx) error {
		query := "UPDATE `transaction` SET `status` = 'approved' WHERE `name` = ?"
		stmt, err := tx.Prepare(query)
		if err != nil {
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(name)
		if err != nil {
			return err
		}
		return nil
	})
}
