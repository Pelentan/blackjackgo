// player.go
package game

type Player struct {
	Name   string
	Bank   int
	Wins   int
	Losses int
}

func NewPlayer(name string) *Player {
	return &Player{
		Name: name,
		Bank: 500, // Start with $500
	}
}
