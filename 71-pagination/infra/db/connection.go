package db

import (
	"ecom/config"
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func GetConnectionString(cnf *config.DBConfig) string {
	connString := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		cnf.User,
		cnf.Password,
		cnf.Host,
		cnf.Port,
		cnf.Name,
		cnf.SslMode,
	)
	return connString
	// return "user=postgres password=admin host=localhost port=5432 dbname=ecom sslmode=disable"
}

func NewConnection(cnf *config.DBConfig) (*sqlx.DB, error) {
	dbSource := GetConnectionString(cnf)

	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return dbCon, nil
}
