package service

import "testing"

func TestPrinter_Print(t *testing.T) {
	tests := []struct {
		name         string
		validatedStr string
		inputStr     string
	}{
		{
			name:         "Case-1",
			validatedStr: "GGGGG",
			inputStr:     "ABCDE",
		},
		{
			name:         "Case-2",
			validatedStr: "GGYYG",
			inputStr:     "ABCDE",
		},
		{
			name:         "Case-3",
			validatedStr: "GGYBB",
			inputStr:     "ABCDE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Printer{}
			p.Print(tt.validatedStr, tt.inputStr)
		})
	}
}
