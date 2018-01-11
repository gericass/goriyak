package local

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func ConnectDB() (*sql.DB, error) {
	dsn := os.Getenv("DSN")
	var cnn *sql.DB
	var err error
	if dsn == "docker" {
		cnn, err = sql.Open("mysql", "root:mysql@tcp(local:3306)/goriyak?parseTime=true")
		if err != nil {
			return nil, err
		}

	} else {
		cnn, err = sql.Open("mysql", "root:mysql@tcp(127.0.0.1:13306)/goriyak?parseTime=true")
		if err != nil {
			return nil, err
		}
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
