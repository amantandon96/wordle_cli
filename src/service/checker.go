package service

import (
	"errors"
	"wordle_cli/src/dao"
)

type IChecker interface {
	Check(correctWord string, inputWord string) (string, error)
}

func NewChecker(wordDao dao.IWordDao) IChecker {
	return &Checker{
		wordDao: wordDao,
	}
}

type Checker struct {
	wordDao dao.IWordDao
}

func (c *Checker) Check(correctWord string, inputWord string) (string, error) {
	if len(inputWord) != len(correctWord) {
		return "", errors.New("input word length doesn't match game's word length")
	}
	if exists := c.wordDao.Exists(inputWord); !exists {
		return "", errors.New("input word is not a valid dictionary word")
	}
	checked := ""
	alphaCount := make(map[rune]int)
	usedAlphaCount := make(map[rune]int)
	for i := 0; i < len(correctWord); i++ {
		alphaCount[rune(correctWord[i])]++
		usedAlphaCount[rune(correctWord[i])] = 0
	}
	for i := 0; i < len(inputWord); i++ {
		usedAlphaCount[rune(inputWord[i])]++
	}
	for i := 0; i < len(inputWord); i++ {
		if inputWord[i] == correctWord[i] {
			checked += "G"
		} else {
			if alphaCount[rune(inputWord[i])] == 0 && usedAlphaCount[rune(inputWord[i])] > 0 {
				checked += "B"
			} else {
				checked += "Y"
			}
		}
	}
	return checked, nil
}
