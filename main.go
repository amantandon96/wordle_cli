package main

import (
	"wordle_cli/config"
	"wordle_cli/flags"
	_ "wordle_cli/flags"

	"wordle_cli/src/service"
)

func main() {
	gamemode := config.V.GetString(flags.Mode)
	length := config.V.GetInt(flags.Length)
	service.GameOrchestrator.Play(length, gamemode)
}

/*
todo:
- makefile to create a binary and install dependencies
- change printErr to logger.Error
- add debug logs
- take wordlength and game mode as command line arguments
- take loglevel as a command line argument, default is error
*/
