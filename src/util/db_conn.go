package util

import (
	"database/sql"
	"fmt"
	//"go.uber.org/zap"
	_ "github.com/lib/pq"
	"wordle_cli/config"
)

var dbConn *sql.DB

func init() {
	dbName := config.V.Get("DB_NAME")
	dbUser := config.V.Get("DB_USER")
	dbHost := config.V.Get("DB_HOST")
	dbPort := config.V.Get("DB_PORT")
	dbPassword := config.V.Get("DB_PASS")
	connStr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	conn, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(fmt.Sprintf("failed to connect to db. name: %v, host: %v, port: %v\n err:%v", dbName, dbHost, dbPort, err))
		//logger.L.With(zap.Error(err)).Fatalf("failed to connect to db. name: %v, host: %v, port: %v", dbName, dbHost, dbPort)
	}
	err = conn.Ping()
	if err != nil {
		//logger.L.With(zap.Error(err)).Fatalf("failed to ping db. name: %v, host: %v, port: %v", dbName, dbHost, dbPort)
		panic(fmt.Sprintf("failed to ping db. name: %v, host: %v, port: %v", dbName, dbHost, dbPort))
	}
	//logger.L.Info("database connected")
	dbConn = conn
}

func GetDBConn() *sql.DB {
	return dbConn
}
