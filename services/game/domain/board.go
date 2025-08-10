package domain

import (
	"errors"
)

// Board - Value Object of game board
type Board struct {
	cells [3][3]string
}

// NewBoard creates a new game board
func NewBoard() *Board {
	return &Board{
		cells: [3][3]string{
			{"", "", ""},
			{"", "", ""},
			{"", "", ""},
		},
	}
}

// MakeMove executes a move on the board
func (b *Board) MakeMove(position int, symbol string) error {
	if position < 0 || position > 8 {
		return errors.New("invalid position")
	}
	
	if symbol != "X" && symbol != "O" {
		return errors.New("invalid symbol")
	}
	
	row := position / 3
	col := position % 3
	
	if b.cells[row][col] != "" {
		return errors.New("position already occupied")
	}
	
	b.cells[row][col] = symbol
	return nil
}

// HasWinner checks if there is a winner
func (b *Board) HasWinner() bool {
	// Check rows
	for i := 0; i < 3; i++ {
		if b.cells[i][0] != "" && b.cells[i][0] == b.cells[i][1] && b.cells[i][1] == b.cells[i][2] {
			return true
		}
	}
	
	// Check columns
	for i := 0; i < 3; i++ {
		if b.cells[0][i] != "" && b.cells[0][i] == b.cells[1][i] && b.cells[1][i] == b.cells[2][i] {
			return true
		}
	}
	
	// Check diagonals
	if b.cells[0][0] != "" && b.cells[0][0] == b.cells[1][1] && b.cells[1][1] == b.cells[2][2] {
		return true
	}
	
	if b.cells[0][2] != "" && b.cells[0][2] == b.cells[1][1] && b.cells[1][1] == b.cells[2][0] {
		return true
	}
	
	return false
}

// IsFull checks if the board is full
func (b *Board) IsFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.cells[i][j] == "" {
				return false
			}
		}
	}
	return true
}

// GetState returns board state as 2D array
func (b *Board) GetState() [][]string {
	result := make([][]string, 3)
	for i := 0; i < 3; i++ {
		result[i] = make([]string, 3)
		for j := 0; j < 3; j++ {
			result[i][j] = b.cells[i][j]
		}
	}
	return result
}

// GetCell returns cell value
func (b *Board) GetCell(row, col int) string {
	if row < 0 || row > 2 || col < 0 || col > 2 {
		return ""
	}
	return b.cells[row][col]
}

// IsValidPosition checks if position is valid
func (b *Board) IsValidPosition(position int) bool {
	if position < 0 || position > 8 {
		return false
	}
	row := position / 3
	col := position % 3
	return b.cells[row][col] == ""
} 