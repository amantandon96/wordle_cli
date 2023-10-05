package service

import (
	"context"
	"fmt"
	"time"
	"wordle_cli/src/dao"
	"wordle_cli/src/models"
	"wordle_cli/src/util"
)

type IGameSvc interface {
	Get(id int) (models.Game, error)
	Start(length int, mode string) (models.Game, error)
	Attempt(attempt models.Attempt) (models.Game, error) // if game is invalid or word is invalid, the attempt is ignored and error is returned else game with it's updated state is returned
	Update(gameId int, state string) error
}

type GameSvc struct {
	wordDao    dao.IWordDao
	gameDao    dao.IGameDao
	attemptDao dao.IGameAttemptDao
}

func (g *GameSvc) Get(id int) (models.Game, error) {
	return g.gameDao.Get(id)
}

func (g *GameSvc) Start(length int, mode string) (models.Game, error) {
	txn, err := util.GetDBConn().BeginTx(context.Background(), nil)
	word, err := g.wordDao.Start(length, mode, txn, false)
	if err != nil {
		txn.Rollback()
		return models.Game{}, err
	}
	game := models.Game{
		Word:      word,
		Mode:      mode,
		State:     models.GameCreated,
		CreatedAt: time.Now(),
	}
	id, err := g.gameDao.Create(game, txn, true)
	if err != nil {
		return models.Game{}, err
	}
	game.Id = id
	return game, nil
}

// gameSvc -> startTxn -> getGame -> err || attempts remaining? -> err|| checkWord -> err || game finished || game failed || game continue -> save attempt -> return
func (g *GameSvc) Attempt(attempt models.Attempt) (models.Game, error) {
	txn, err := util.GetDBConn().BeginTx(context.Background(), nil)
	if err != nil {
		return models.Game{}, err
	}
	game, err := g.gameDao.Get(attempt.GameId, txn, false)
	if err != nil {
		return game, err
	}
	if len(game.Attempts) == game.MaxAttempts || game.State == models.GameFailed {
		return game, fmt.Errorf("game max attempts completed")
	}
	if !g.wordDao.Exists(attempt.Word) {
		return game, fmt.Errorf("word %v does not exist", attempt.Word)
	}

	err = g.attemptDao.CreateAttempt(attempt, txn, true)
	if err != nil {
		return game, err
	}
	game.Attempts = append(game.Attempts, attempt)
	return game, nil
}

func (g *GameSvc) Update(gameId int, state string) error {
	return g.gameDao.Update(gameId, state)
}
func NewGameSvc(wordDao dao.IWordDao, gameDao dao.IGameDao, attemptDao dao.IGameAttemptDao) IGameSvc {
	return &GameSvc{
		wordDao:    wordDao,
		gameDao:    gameDao,
		attemptDao: attemptDao,
	}
}
