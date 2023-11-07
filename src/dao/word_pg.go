package dao

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type WordDaoPg struct {
	pgConn *sql.DB
}

func (w *WordDaoPg) Start(length int, mode string, varargs ...interface{}) (string, error) {
	txn, commit := parseTxnArgs(varargs)
	if txn == nil {
		t, err := w.pgConn.BeginTx(context.Background(), nil)
		txn = t
		if err != nil {
			return "", err
		}
	}
	query := `select w.word from words as w left outer join game as g on w.word=g.word where w.length=$1 order by coalesce(g.created_at, '1970-12-01 00:00:00'),random() limit 1`
	row := txn.QueryRow(query, length)
	var word string
	err := row.Scan(&word)

	if err != nil {
		err2 := txn.Rollback()
		if err2 != nil {
			return "", fmt.Errorf("error %v occurred followed by this error: %v", err, err2.Error())
		}
		return "", err
	}
	if commit {
		err = txn.Commit()
		if err != nil {
			return "", err
		}
	}
	return word, nil
}

func (w *WordDaoPg) Exists(word string) bool {

	query := `select word from words where word=$1`
	row := w.pgConn.QueryRow(query, word)
	var queriedWord string
	_ = row.Scan(&queriedWord)
	return queriedWord == word
}

func parseTxnArgs(varargs ...interface{}) (*sql.Tx, bool) {
	var txn *sql.Tx
	var commit bool = true
	if len(varargs) > 1 {
		if t, ok := varargs[0].(*sql.Tx); ok {
			txn = t
		}
		if c, ok := varargs[1].(bool); ok {
			commit = c
		}

	}
	return txn, commit
}
