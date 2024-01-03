package main

type Player struct {
	ID          int
	Name        string
	Score       int
	WordToGuess []string
	IsGuessing  bool
}

func (p *Player) SetName(name string) {
	p.Name = name
}

func (p *Player) SetScore(score int) {
	p.Score = score
}

func (p *Player) SetIsGuessing(isGuessing bool) {
	p.IsGuessing = isGuessing
}

func (p *Player) AddWordToGuess(guess string) {
	p.WordToGuess = append(p.WordToGuess, guess)
}
