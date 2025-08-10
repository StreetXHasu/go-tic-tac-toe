package domain

import (
	"errors"
	"time"
)

// Move - Value Object of a move
type Move struct {
	Position   int
	Symbol     string
	PlayerID   string
	Timestamp  time.Time
}

// NewMove creates a new move
func NewMove(position int, symbol, playerID string) (*Move, error) {
	if err := validateMove(position, symbol); err != nil {
		return nil, err
	}
	
	return &Move{
		Position:  position,
		Symbol:    symbol,
		PlayerID:  playerID,
		Timestamp: time.Now(),
	}, nil
}

// validateMove validates move
func validateMove(position int, symbol string) error {
	if position < 0 || position > 8 {
		return errors.New("position must be between 0 and 8")
	}
	
	if symbol != "X" && symbol != "O" {
		return errors.New("symbol must be X or O")
	}
	
	return nil
}

// IsValid checks if move is valid
func (m *Move) IsValid() bool {
	return m.Position >= 0 && m.Position <= 8 && 
		   (m.Symbol == "X" || m.Symbol == "O") &&
		   m.PlayerID != ""
} 