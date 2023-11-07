package service

import (
	"testing"
	"wordle_cli/src/dao/mock_dao"
)

func TestChecker_Check(t *testing.T) {
	dao := mock_dao.MockWordDao{}
	type args struct {
		correctWord string
		inputWord   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "unique alpha-wrong",
			args: args{
				correctWord: "ABCDE",
				inputWord:   "PQRST",
			},
			want:    "BBBBB",
			wantErr: false,
		},
		{
			name: "unique alpha - mix",
			args: args{
				correctWord: "ABCDE",
				inputWord:   "EFGAB",
			},
			want:    "YBBYY",
			wantErr: false,
		},
		{
			name: "",
			args: args{
				correctWord: "ABCDE",
				inputWord:   "EFGAB",
			},
			want:    "YBBYY",
			wantErr: false,
		},
		{
			name: "",
			args: args{
				correctWord: "ABCDD",
				inputWord:   "ADDKK",
			},
			want:    "GYYBB",
			wantErr: false,
		},
		{
			name: "",
			args: args{
				correctWord: "ABCDD",
				inputWord:   "AKDDK",
			},
			want:    "GBYGB",
			wantErr: false,
		},
		{
			name: "",
			args: args{
				correctWord: "ABCDD",
				inputWord:   "AKKDD",
			},
			want:    "GBBGG",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Checker{
				wordDao: &dao,
			}
			got, err := c.Check(tt.args.correctWord, tt.args.inputWord)
			if (err != nil) != tt.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Check() got = %v, want %v", got, tt.want)
			}
		})
	}
}
