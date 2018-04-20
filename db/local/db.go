package local

import (
	"database/sql"
	"os"

	"github.com/gericass/goriyak/setting"
	_ "github.com/go-sql-driver/mysql"
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
		user := setting.ServerConfig.Mysql.User
		password := setting.ServerConfig.Mysql.Password
		host := setting.ServerConfig.Mysql.Host
		port := setting.ServerConfig.Mysql.Port
		cnn, err = sql.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/goriyak?parseTime=true")
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
