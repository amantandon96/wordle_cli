package main

import "wordle_cli/src/service"

func main() {
	//for i := 40; i < 50; i++ {
	//	fmt.Println(fmt.Sprintf("\x1b[%dm%s\x1b[0m", i, "Hello"))
	//}
	//tests := []struct {
	//	name         string
	//	validatedStr string
	//	inputStr     string
	//}{
	//	{
	//		name:         "Case-1",
	//		validatedStr: "GGGGG",
	//		inputStr:     "ABCDE",
	//	},
	//	{
	//		name:         "Case-2",
	//		validatedStr: "GGYYG",
	//		inputStr:     "ABCDE",
	//	},
	//	{
	//		name:         "Case-3",
	//		validatedStr: "GGYBB",
	//		inputStr:     "ABCDE",
	//	},
	//}
	//for _, tt := range tests {
	//
	//	p := &service.Printer{}
	//	p.Print(tt.validatedStr, tt.inputStr)
	//
	//}
	//service.StartGame(5)
	service.GameOrchestrator.Play(5, "practice")
}
