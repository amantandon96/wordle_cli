package util

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"

	//"go.uber.org/zap"
	_ "github.com/mattn/go-sqlite3"
	"wordle_cli/config"
)

var dbConn *sql.DB

const bootstrapScriptPath = "resources/bootstrap_db.py"

func init() {
	parameterizedDbPath := config.V.GetString("SQLITE_DB_PATH")
	userName, err := getUser()
	if err != nil {
		panic(fmt.Sprintf("failed to get current user. error: %v", err))
	}
	dbPath := strings.ReplaceAll(parameterizedDbPath, "${current_user}", userName)
	_, err = os.Stat(dbPath)
	if os.IsNotExist(err) {
		bootstrapDb(dbPath)
	}
	conn, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to db. dir:%v \n err:%v", dbPath, err))
	}
	err = conn.Ping()
	if err != nil {
		panic(fmt.Sprintf("failed to ping db. path: %v", dbPath))
	}
	dbConn = conn
}

func GetDBConn() *sql.DB {
	return dbConn
}

func bootstrapDb(dbPath string) {

	cmd := exec.Command("python3", bootstrapScriptPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(fmt.Sprintf("Error running bootstrap script: %v\n", err))
	}
}
func getUser() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	return currentUser.Username, nil
}
