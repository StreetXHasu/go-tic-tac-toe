package domain

import (
	"errors"
	"time"
)

// Game - main game aggregate
type Game struct {
	ID          string
	Player1     *Player
	Player2     *Player
	Board       *Board
	Status      GameStatus
	CurrentTurn *Player
	Winner      *Player
	CreatedAt   time.Time
	StartedAt   *time.Time
	FinishedAt  *time.Time
}

// GameStatus - game status
type GameStatus string

const (
	GameStatusWaiting GameStatus = "waiting"
	GameStatusActive  GameStatus = "active"
	GameStatusFinished GameStatus = "finished"
)

// NewGame creates a new game
func NewGame(player1 *Player) *Game {
	return &Game{
		ID:        generateGameID(),
		Player1:   player1,
		Board:     NewBoard(),
		Status:    GameStatusWaiting,
		CreatedAt: time.Now(),
	}
}

// JoinGame allows the second player to join the game
func (g *Game) JoinGame(player2 *Player) error {
	if g.Status != GameStatusWaiting {
		return errors.New("game is not in waiting status")
	}
	
	if g.Player1.ID == player2.ID {
		return errors.New("player cannot join their own game")
	}
	
	g.Player2 = player2
	g.Status = GameStatusActive
	g.CurrentTurn = g.Player1
	now := time.Now()
	g.StartedAt = &now
	
	return nil
}

// MakeMove executes a move in the game
func (g *Game) MakeMove(player *Player, position int) error {
	if g.Status != GameStatusActive {
		return errors.New("game is not active")
	}
	
	if g.CurrentTurn.ID != player.ID {
		return errors.New("not your turn")
	}
	
	if !g.isPlayerInGame(player) {
		return errors.New("player is not in this game")
	}
	
	if err := g.Board.MakeMove(position, player.Symbol); err != nil {
		return err
	}
	
	// Check for win
	if g.Board.HasWinner() {
		g.Status = GameStatusFinished
		g.Winner = player
		now := time.Now()
		g.FinishedAt = &now
		return nil
	}
	
	// Check for draw
	if g.Board.IsFull() {
		g.Status = GameStatusFinished
		now := time.Now()
		g.FinishedAt = &now
		return nil
	}
	
	// Switch turn to other player
	g.switchTurn()
	
	return nil
}

// isPlayerInGame checks if player is a participant in the game
func (g *Game) isPlayerInGame(player *Player) bool {
	return (g.Player1 != nil && g.Player1.ID == player.ID) ||
		   (g.Player2 != nil && g.Player2.ID == player.ID)
}

// switchTurn switches turn between players
func (g *Game) switchTurn() {
	if g.CurrentTurn.ID == g.Player1.ID {
		g.CurrentTurn = g.Player2
	} else {
		g.CurrentTurn = g.Player1
	}
}

// GetGameState returns current game state
func (g *Game) GetGameState() GameState {
	return GameState{
		ID:          g.ID,
		Status:      g.Status,
		Board:       g.Board.GetState(),
		CurrentTurn: g.CurrentTurn,
		Winner:      g.Winner,
		Player1:     g.Player1,
		Player2:     g.Player2,
	}
}

// GameState - game state for client transmission
type GameState struct {
	ID          string     `json:"id"`
	Status      GameStatus `json:"status"`
	Board       [][]string `json:"board"`
	CurrentTurn *Player    `json:"current_turn"`
	Winner      *Player    `json:"winner,omitempty"`
	Player1     *Player    `json:"player1"`
	Player2     *Player    `json:"player2,omitempty"`
}

// generateGameID generates unique game ID
func generateGameID() string {
	return "game_" + time.Now().Format("20060102150405")
} 