package main

import (
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
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Doğru harf doğru konum - yeşil renk",
			fields: fields{
				Letter: 'A',
				Color:  "",
			},
			args: args{
				color: "green",
			},
			want: "green",
		},
		{
			name: "Doğru harf yanlış konum - sarı renk",
			fields: fields{
				Letter: 'A',
				Color:  "",
			},
			args: args{
				color: "yellow",
			},
			want: "yellow",
		},
		{
			name: "Yanlış harf - gri renk",
			fields: fields{
				Letter: 'A',
				Color:  "",
			},
			args: args{
				color: "gray",
			},
			want: "gray",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Alphabet{
				Letter: tt.fields.Letter,
				Color:  tt.fields.Color,
			}
			a.SetColor(tt.args.color)
			if a.Color != tt.want {
				t.Errorf("SetColor() = %v, want %v", a.Color, tt.want)
			}
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
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Doğru harf doğru konum - yeşil renk",
			fields: fields{
				Letter: 'A',
				Color:  "",
			},
			args: args{
				color: "green",
			},
			want: "green",
		},
		{
			name: "Doğru harf yanlış konum - sarı renk",
			fields: fields{
				Letter: 'A',
				Color:  "",
			},
			args: args{
				color: "yellow",
			},
			want: "yellow",
		},
		{
			name: "Yanlış harf - gri renk",
			fields: fields{
				Letter: 'A',
				Color:  "",
			},
			args: args{
				color: "gray",
			},
			want: "gray",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Forecast{
				Letter: tt.fields.Letter,
				Color:  tt.fields.Color,
			}
			f.SetColor(tt.args.color)
			if f.Color != tt.want {
				t.Errorf("SetColor() = %v, want %v", f.Color, tt.want)
			}
		})
	}
}

func TestWordle_SetAlphabet(t *testing.T) {
	type fields struct {
		Word      string
		Forecasts []Forecasts
		Alphabet  []Alphabet
	}
	type args struct {
		letter rune
		color  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Alfabede harf rengi güncelleme - yeşil",
			fields: fields{
				Word: "KALEM",
				Alphabet: []Alphabet{
					{Letter: 'K', Color: ""},
					{Letter: 'A', Color: ""},
					{Letter: 'L', Color: ""},
					{Letter: 'E', Color: ""},
					{Letter: 'M', Color: ""},
				},
			},
			args: args{
				letter: 'K',
				color:  "green",
			},
			want: "green",
		},
		{
			name: "Alfabede harf rengi güncelleme - sarı",
			fields: fields{
				Word: "KALEM",
				Alphabet: []Alphabet{
					{Letter: 'K', Color: ""},
					{Letter: 'A', Color: ""},
					{Letter: 'L', Color: ""},
					{Letter: 'E', Color: ""},
					{Letter: 'M', Color: ""},
				},
			},
			args: args{
				letter: 'A',
				color:  "yellow",
			},
			want: "yellow",
		},
		{
			name: "Alfabede harf rengi güncelleme - gri",
			fields: fields{
				Word: "KALEM",
				Alphabet: []Alphabet{
					{Letter: 'K', Color: ""},
					{Letter: 'A', Color: ""},
					{Letter: 'L', Color: ""},
					{Letter: 'E', Color: ""},
					{Letter: 'M', Color: ""},
				},
			},
			args: args{
				letter: 'Z',
				color:  "gray",
			},
			want: "gray",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Wordle{
				Word:      tt.fields.Word,
				Forecasts: tt.fields.Forecasts,
				Alphabet:  tt.fields.Alphabet,
			}
			w.SetAlphabet(tt.args.letter, tt.args.color)
			for _, a := range w.Alphabet {
				if a.Letter == tt.args.letter && a.Color != tt.want {
					t.Errorf("SetAlphabet() = %v, want %v", a.Color, tt.want)
				}
			}
		})
	}
}
