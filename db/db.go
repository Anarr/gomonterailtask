package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DbInterface interface {
	Connect() (*sql.DB, error)
}

//MysqlDB store credentials for mysql database connection
type MysqlDb struct {
	Host, Username, Pass, Name, Client string
	Port                               int
}

//Connect create db connection
func (md *MysqlDb) Connect() (*sql.DB, error)  {
	connStr := fmt.Sprintf("%v:%v@tcp(%v:%d)/%v", md.Username, md.Pass, md.Host, md.Port, md.Name)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

//New retrieve new database connection instance
func New(db DbInterface) (*sql.DB, error) {
	return db.Connect()
}