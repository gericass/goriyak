package local

import (
	"database/sql"
)

func ScanTransactions(rows *sql.Rows) ([]*LocalTransaction, error) {
	defer rows.Close()
	transactions := make([]*LocalTransaction, 0)

	for rows.Next() {
		t := new(LocalTransaction)
		if err := rows.Scan(&t.ID, &t.Name, &t.SendNodeID, t.ReceiveNodeID, &t.Amount, &t.CreatedAt); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)

	}
	return transactions, nil
}
