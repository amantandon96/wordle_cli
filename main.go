package main

import "wordle_cli/src/service"

func main() {
	service.GameOrchestrator.Play(5, "practice")
}

/*
todo:
- makefile to create a binary and install dependencies
- change printErr to logger.Error
- add debug logs
- take wordlength and game mode as command line arguments
- take loglevel as a command line argument, default is error
*/
