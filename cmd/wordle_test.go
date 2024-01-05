package main

import (
	"strconv"
	"testing"
)

func TestAlphabet_SetColor(t *testing.T) {
	type fields struct {
		Letter rune
		Color  string
	}
	type args struct {
		color string
	}
	tests := []struct {
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			a := &Alphabet{
				Letter: tt.fields.Letter,
				Color:  tt.fields.Color,
			}
			a.SetColor(tt.args.color)
		})
	}
}

func TestForecast_SetColor(t *testing.T) {
	type fields struct {
		Letter rune
		Color  string
	}
	type args struct {
		color string
	}
	tests := []struct {
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			f := &Forecast{
				Letter: tt.fields.Letter,
				Color:  tt.fields.Color,
			}
			f.SetColor(tt.args.color)
		})
	}
}

func TestWordle_CheckWord(t *testing.T) {
	type fields struct {
		Word      string
		Forecasts map[string][]Forecast
		Alphabet  []Alphabet
	}
	type args struct {
		word string
	}
	tests := []struct {
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			w := &Wordle{
				Word:      tt.fields.Word,
				Forecasts: tt.fields.Forecasts,
				Alphabet:  tt.fields.Alphabet,
			}
			w.CheckWord(tt.args.word)
		})
	}
}

func TestWordle_SetAlphabet(t *testing.T) {
	type fields struct {
		Word      string
		Forecasts map[string][]Forecast
		Alphabet  []Alphabet
	}
	type args struct {
		letter rune
		color  string
	}
	tests := []struct {
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			w := &Wordle{
				Word:      tt.fields.Word,
				Forecasts: tt.fields.Forecasts,
				Alphabet:  tt.fields.Alphabet,
			}
			w.SetAlphabet(tt.args.letter, tt.args.color)
		})
	}
}
