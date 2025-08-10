package domain

import (
	"time"
)

// Player - player entity
type Player struct {
	ID        string
	Username  string
	Email     string
	Symbol    string // "X" or "O"
	CreatedAt time.Time
}

// NewPlayer creates a new player
func NewPlayer(id, username, email string) *Player {
	return &Player{
		ID:        id,
		Username:  username,
		Email:     email,
		CreatedAt: time.Now(),
	}
}

// AssignSymbol assigns symbol to player
func (p *Player) AssignSymbol(symbol string) {
	p.Symbol = symbol
}

// IsValidSymbol checks if symbol is valid
func (p *Player) IsValidSymbol() bool {
	return p.Symbol == "X" || p.Symbol == "O"
} 