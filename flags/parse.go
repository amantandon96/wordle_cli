package flags

import (
	"flag"
	"wordle_cli/config"
)

func init() {
	debug := flag.Bool("debug", false, "to enable debug logs")
	length := flag.Int("length", 5, "length of word: a number between 5 and 25")
	mode := flag.String("mode", "practice", "game mode: [practice, daily]")
	flag.Parse()
	config.V.Set(Debug, *debug)
	config.V.Set(Length, *length)
	config.V.Set(Mode, *mode)
}
