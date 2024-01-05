package main

import (
	"os"
	"reflect"
	"strconv"
	"testing"
)

func TestContainsValueInSlice(t *testing.T) {
	type args struct {
		val   any
		array any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := ContainsValueInSlice(tt.args.val, tt.args.array); got != tt.want {
				t.Errorf("ContainsValueInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExistsLetter(t *testing.T) {
	type args struct {
		word   string
		letter byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := ExistsLetter(tt.args.word, tt.args.letter); got != tt.want {
				t.Errorf("ExistsLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLineCount(t *testing.T) {
	type args struct {
		file *os.File
	}
	tests := []struct {
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := GetLineCount(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLineCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetLineCount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetWordFromFile(t *testing.T) {
	type args struct {
		file   *os.File
		lineNo int
	}
	tests := []struct {
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := GetWordFromFile(tt.args.file, tt.args.lineNo)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWordFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetWordFromFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetWords(t *testing.T) {
	type args struct {
		lang   string
		length int
	}
	tests := []struct {
		args    args
		want    string
		wantErr bool
	}{
		{args: args{lang: "en", length: 5}, want: "", wantErr: false},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := GetWords(tt.args.lang, tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetWords() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandleLog(t *testing.T) {
	type args struct {
		message string
		err     error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			HandleLog(tt.args.message, tt.args.err)
		})
	}
}

func TestRGB(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := RGB(); got != tt.want {
				t.Errorf("RGB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomColor(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := RandomColor(); got != tt.want {
				t.Errorf("RandomColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomName(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := RandomName(tt.args.n); got != tt.want {
				t.Errorf("RandomName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetAlphabet(t *testing.T) {
	type args struct {
		lang string
	}
	tests := []struct {
		name string
		args args
		want []Alphabet
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := SetAlphabet(tt.args.lang); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetAlphabet() = %v, want %v", got, tt.want)
			}
		})
	}
}
