package service

import (
	"wordle_cli/src/dao"
)

type Word interface {
	Get(length int) (string, error)
}

type WordSvc struct {
	wordDao dao.IWordDao
}

const defaultSelector = "random"

func (w *WordSvc) Get(length int) (string, error) {
	return w.wordDao.Start(length, defaultSelector)
}
