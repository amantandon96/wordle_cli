package dao

import (
	"context"
	"database/sql"
	"fmt"
	"wordle_cli/src/models"
	"wordle_cli/src/util"
)

type IGameDao interface {
	Create(game models.Game, varargs ...interface{}) (int, error)
	Get(id int, varargs ...interface{}) (models.Game, error)
	Update(gameId int, state string, varargs ...interface{}) error
}

type IGameAttemptDao interface {
	CreateAttempt(attempt models.Attempt, varargs ...interface{}) error
}

type GameDaoPg struct {
	pgConn *sql.DB
}

func NewGameDao() IGameDao {
	return &GameDaoPg{util.GetDBConn()}
}

func NewGameAttemptDao() IGameAttemptDao {
	return &GameDaoPg{util.GetDBConn()}
}

func (g *GameDaoPg) CreateAttempt(attempt models.Attempt, varargs ...interface{}) error {
	txn, commit := parseTxnArgs(varargs)
	if txn == nil {
		t, err := g.pgConn.BeginTx(context.Background(), nil)
		txn = t
		if err != nil {
			return err
		}
	}
	insertQ := `insert into game_log (game_id, attempted_word, attempt_num) VALUES ($1, $2, $3)`
	_, err := txn.Exec(insertQ, attempt.GameId, attempt.Word, attempt.AttemptNum)
	if err != nil {
		err2 := txn.Rollback()
		if err2 != nil {
			return fmt.Errorf("error %v occurred followed by this error: %v", err, err2.Error())
		}
		return err
	}
	if commit {
		err = txn.Commit()
		if err != nil {
			return err
		}
	}
	return nil
}
func (g *GameDaoPg) Create(game models.Game, varargs ...interface{}) (int, error) {
	txn, commit := parseTxnArgs(varargs)
	if txn == nil {
		t, err := g.pgConn.BeginTx(context.Background(), nil)
		txn = t
		if err != nil {
			return 0, err
		}
	}
	insertQ := `insert into game (word, mode, state) values ($1, $2, $3) returning id`

	row := txn.QueryRow(insertQ, game.Word, game.Mode, game.State)
	var id int
	err := row.Scan(&id)
	if err != nil {
		err2 := txn.Rollback()
		if err2 != nil {
			return 0, fmt.Errorf("error %v occurred followed by this error: %v", err, err2.Error())
		}
		return 0, err
	}
	if commit {
		err = txn.Commit()
		if err != nil {
			return 0, err
		}
	}
	return id, nil
}

func (g *GameDaoPg) Get(id int, varargs ...interface{}) (models.Game, error) {
	txn, commit := parseTxnArgs(varargs)
	if txn == nil {
		t, err := g.pgConn.BeginTx(context.Background(), nil)
		txn = t
		if err != nil {
			return models.Game{}, err
		}
	}
	query := `select g.id as id, g.word as word, g.mode as mode, g.state as state, g.created_at as created_at, a.attempt_num as attempt_num, a.attempted_word as attempted_word, att.max_attempts as max_attempts 
from game as g left join game_log as a on g.id=a.game_id left join attempts as att on length(g.word)=att.word_len where id=$1`
	rows, err := txn.Query(query, id)
	if err != nil {
		return models.Game{}, err
	}
	if rows.Err() != nil {
		return models.Game{}, rows.Err()
	}
	var attempts []models.Attempt
	var game models.Game
	for rows.Next() {
		var attempt models.Attempt
		var attemptedWord sql.NullString
		var numAttempt sql.NullInt64
		err = rows.Scan(&game.Id, &game.Word, &game.Mode, &game.State, &game.CreatedAt, &numAttempt, &attemptedWord, &game.MaxAttempts)
		if err != nil {
			return game, err
		}
		attempt.AttemptNum = int(numAttempt.Int64)
		attempt.Word = attemptedWord.String
		if attempt.Word != "" {
			attempts = append(attempts, attempt)
		}
	}
	game.Attempts = attempts
	if err != nil {
		err2 := txn.Rollback()
		if err2 != nil {
			return models.Game{}, fmt.Errorf("error %v occurred followed by this error: %v", err, err2.Error())
		}
		return models.Game{}, err
	}
	if commit {
		err = txn.Commit()
		if err != nil {
			return models.Game{}, err
		}
	}
	return game, nil
}

func (g *GameDaoPg) Update(gameId int, state string, varargs ...interface{}) error {
	txn, commit := parseTxnArgs(varargs)
	if txn == nil {
		t, err := g.pgConn.BeginTx(context.Background(), nil)
		txn = t
		if err != nil {
			return err
		}
	}
	updateQ := `update game set state=$2 where id=$1`
	_, err := txn.Exec(updateQ, gameId, state)
	if err != nil {
		err2 := txn.Rollback()
		if err2 != nil {
			return fmt.Errorf("error %v occurred followed by this error: %v", err, err2.Error())
		}
		return err
	}
	if commit {
		err = txn.Commit()
		if err != nil {
			return err
		}
	}
	return nil

}

// flow for start game:
// gameSvc->startTxn -> getWord -> create game -> commit txn
// flow for an attempt:
// gameSvc -> startTxn -> getGame -> err || attempts remaining? -> err|| checkWord -> err || game finished || game failed || game continue -> save attempt -> return
