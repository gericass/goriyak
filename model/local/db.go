package local

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func ConnectDB() (*sql.DB, error) {
	cnn, err := sql.Open("mysql", "root:mysql@tcp(127.0.0.1:13306)/goriyak?parseTime=true")
	if err != nil {
		return nil, err
	}
	return cnn, nil
}

func dbTransaction(tx *sql.Tx, s func(*sql.Tx) error) error {
	var err error

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = s(tx)
	return err
}
