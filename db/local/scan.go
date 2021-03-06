package local

import (
	"database/sql"
)

func scanTransactions(rows *sql.Rows) ([]*LocalTransaction, error) {
	defer rows.Close()
	transactions := make([]*LocalTransaction, 0)

	for rows.Next() {
		t := new(LocalTransaction)
		if err := rows.Scan(&t.ID, &t.Name, &t.SendNodeID, &t.ReceiveNodeID, &t.Amount, &t.Status, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)

	}
	return transactions, nil
}

func scanCount(row *sql.Row) (int, error) {
	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}
