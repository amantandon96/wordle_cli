package dao

import (
	"sync"
	"wordle_cli/src"
	"wordle_cli/src/util"
)

type IWordDao interface {
	Start(length int, mode string, varargs ...interface{}) (string, error)
	Exists(word string) bool
}

type WordLocalDao struct {
	sync.Mutex
	wordList         map[int][]string
	selectorRegistry src.ISelectorRegistry
}

func NewWordDao(daotype string, wordList []string, registry src.ISelectorRegistry) IWordDao {
	switch daotype {
	case "local":
		{
			wdMap := make(map[int][]string)
			wdMap[5] = wordList
			return &WordLocalDao{
				wordList:         wdMap,
				selectorRegistry: registry,
			}
		}
	case "remote":
		{
			return &WordDaoPg{pgConn: util.GetDBConn()}
		}
	default:
		panic("invalid case")
	}
}

func (w *WordLocalDao) Start(length int, mode string, varargs ...interface{}) (string, error) {
	return "worry", nil
}

func (w *WordLocalDao) Exists(word string) bool {
	return true
}

// tables:
// word table -> length, word
// attempt -> word_len, max_attempts
// games -> id, word, game_mode, state, created_at
// game_log -> game_id, attempted_word, attempt_num
//
// selection_logic: least_freq_used
// left-inner-join word with games, filter where game.game_mode='daily', sort by created_at, pick the first one. Left join will ensure that word with null created_at gets picked first.
// but this will give alphabetical results. Shuffle order before joining
