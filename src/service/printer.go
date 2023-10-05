package service

import "fmt"

type IPrinter interface {
	Print(validatedStr string, inputString string)
	PrintErr(error)
}

type Printer struct{}

func NewPrinter() IPrinter {
	return &Printer{}
}

func (p *Printer) Print(validatedStr string, inputString string) {
	output := ""
	for i := 0; i < len(validatedStr); i++ {
		switch validatedStr[i : i+1] {
		case "G":
			output += Green(inputString[i : i+1])
		case "Y":
			output += Yellow(inputString[i : i+1])
		case "B":
			output += Red(inputString[i : i+1])
		}
	}
	fmt.Println(output)
	return
}

func (p *Printer) PrintErr(err error) {
	fmt.Println(fmt.Sprintf("Error occurred: %v", err.Error()))
}
func Red(str string) string {
	return fmt.Sprintf("\x1b[01;04;%dm%s\x1b[0m", 41, str)
}
func Green(str string) string {
	return fmt.Sprintf("\x1b[01;04;%dm%s\x1b[0m", 42, str)
}
func Yellow(str string) string {
	return fmt.Sprintf("\x1b[01;04;%dm%s\x1b[0m", 43, str)
}
