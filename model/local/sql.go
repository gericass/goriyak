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
func GetTransactionsByTime(start time.Time, end time.Time, db *sql.DB) ([]*LocalTransaction, error) {

	query := "SELECT `id`,`name`,`send_node_id`,`receive_node_id`,`amount`,`status`,`created_at`,`updated_at` FROM `transaction` WHERE ? <= `updated_at` AND `updated_at` < ? AND `status` = 'approved'"
	rows, err := db.Query(query, start, end)
	if err != nil {
		return nil, err
	}
	return scanTransactions(rows)
}

// GetTransactionsByName : to get transactions by name column
func GetTransactionsByName(name string, db *sql.DB) ([]*LocalTransaction, error) {

	query := "SELECT `id`,`name`,`send_node_id`,`receive_node_id`,`amount`,`status`,`created_at`,`updated_at` FROM `transaction` WHERE `name` = ? LIMIT 2"
	rows, err := db.Query(query, name)
	if err != nil {
		return nil, err
	}
	return scanTransactions(rows)
}

// GetTransactionExists : to search transactions by name column
func GetTransactionExists(name string, db *sql.DB) (bool, error) {
	query := "SELECT count(*) FROM `transaction` WHERE `name` = ? LIMIT 2"
	row := db.QueryRow(query, name)
	count, err := scanCount(row)
	if err != nil {
		return false, err
	}
	if count < 1 {
		return false, nil
	}
	return true, nil
}

// PutTransaction : to put transactions to MySQL transaction table
func (t *LocalTransaction) PutTransaction(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
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
func DeleteTransactionByTime(time time.Time, db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
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
func UpdateTransactionStatus(name string, currentTime time.Time, db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	return dbTransaction(tx, func(tx *sql.Tx) error {
		query := "UPDATE `transaction` SET `status` = 'approved', `updated_at` = ? WHERE `name` = ?"
		stmt, err := tx.Prepare(query)
		if err != nil {
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(name, currentTime)
		if err != nil {
			return err
		}
		return nil
	})
}
