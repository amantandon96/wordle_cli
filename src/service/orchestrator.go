package service

import (
	"bufio"
	"fmt"
	"os"
	"wordle_cli/src/dao"
	"wordle_cli/src/models"
)

var GameOrchestrator GameOrch

func init() {
	wordDao := dao.NewWordDao("remote", []string{}, nil)
	checker := NewChecker(wordDao)
	gameDao := dao.NewGameDao()
	attemptDao := dao.NewGameAttemptDao()
	gameService := NewGameSvc(wordDao, gameDao, attemptDao)
	printer := NewPrinter()
	GameOrchestrator = GameOrch{
		gameSvc: gameService,
		checker: checker,
		printer: printer,
	}
}

type GameOrch struct {
	gameSvc IGameSvc
	checker IChecker
	printer IPrinter
}

func (orch *GameOrch) Play(wordLength int, mode string) {
	if wordLength > 25 || mode != "practice" {
		orch.printer.PrintErr(fmt.Errorf("invalid input"))
		return
	}
	game, err := orch.gameSvc.Start(wordLength, mode)
	if err != nil {
		orch.printer.PrintErr(err)
		return
	}
	game, err = orch.gameSvc.Get(game.Id)
	if err != nil {
		orch.printer.PrintErr(err)
		return
	}
	numAttempt := len(game.Attempts) + 1
	for numAttempt <= game.MaxAttempts {

		inputW := GetInput()
		checked, err := orch.checker.Check(game.Word, inputW)
		_, err = orch.gameSvc.Attempt(models.Attempt{Word: inputW, AttemptNum: numAttempt, GameId: game.Id})
		if err != nil {
			orch.printer.PrintErr(err)
			continue
		}
		numAttempt++
		orch.printer.Print(checked, inputW)
		gameWon := true
		for i := 0; i < wordLength; i++ {
			if checked[i:i+1] != "G" {
				gameWon = false
				break
			}
		}
		if gameWon {
			err = orch.gameSvc.Update(game.Id, models.GameCompletedSuccessfully)
			if err != nil {
				orch.printer.PrintErr(err)
				return
			}
			fmt.Println("You have won the game")
			break
		}

		if numAttempt == game.MaxAttempts+1 {
			err = orch.gameSvc.Update(game.Id, models.GameFailed)
			if err != nil {
				orch.printer.PrintErr(err)
				return
			}
			fmt.Println("You have lost the game")
			fmt.Println("The correct word is: ", game.Word)
		}
	}

}

//
//func StartGame(wordLength int) {
//	wordDao := dao.NewWordDao([]string{"words", "abcde", "check"}, nil)
//	wordSvc := WordSvc{wordDao}
//	checker := NewChecker(wordDao)
//	printer := Printer{}
//	word, err := wordSvc.Get(wordLength)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(fmt.Sprintf("Enter a %d alphabet word", wordLength))
//	numAttempts := 1 + wordLength
//	for numAttempts > 0 {
//		inputW := GetInput()
//		checked, err := checker.Check(word, inputW)
//		if err != nil {
//			fmt.Println(err)
//			continue
//		}
//		printer.Print(checked, inputW)
//		gameWon := true
//		for i := 0; i < wordLength; i++ {
//			if checked[i:i+1] != "G" {
//				gameWon = false
//				break
//			}
//		}
//		if gameWon {
//			fmt.Println("You have won the game")
//			break
//		}
//		numAttempts--
//		if numAttempts == 0 {
//			fmt.Println("You have lost the game")
//			fmt.Println("The correct word is: ", word)
//		}
//	}
//}

func GetInput() string {
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')

	// Remove the newline character from the input
	userInput = userInput[:len(userInput)-1]

	return userInput
}
