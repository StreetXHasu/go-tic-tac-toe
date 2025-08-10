package domain

import (
	"errors"
	"math/rand"
	"time"
)

// GameService - domain service for game logic
type GameService struct{}

// NewGameService creates a new game service
func NewGameService() *GameService {
	return &GameService{}
}

// CreateGame creates a new game
func (gs *GameService) CreateGame(player1 *Player) *Game {
	// Assign symbols to players
	player1.AssignSymbol("X")
	
	game := NewGame(player1)
	return game
}

// JoinGame allows player to join the game
func (gs *GameService) JoinGame(game *Game, player2 *Player) error {
	if game.Player2 != nil {
		return errors.New("game is already full")
	}
	
	// Assign symbol to second player
	player2.AssignSymbol("O")
	
	return game.JoinGame(player2)
}

// MakeMove executes a move in the game
func (gs *GameService) MakeMove(game *Game, player *Player, position int) error {
	// Create move
	move, err := NewMove(position, player.Symbol, player.ID)
	if err != nil {
		return err
	}
	
	// Execute move
	return game.MakeMove(player, move.Position)
}

// GetAvailableGames returns list of available games
func (gs *GameService) GetAvailableGames(games []*Game) []*Game {
	var available []*Game
	for _, game := range games {
		if game.Status == GameStatusWaiting {
			available = append(available, game)
		}
	}
	return available
}

// GetPlayerGames возвращает игры игрока
func (gs *GameService) GetPlayerGames(games []*Game, playerID string) []*Game {
	var playerGames []*Game
	for _, game := range games {
		if (game.Player1 != nil && game.Player1.ID == playerID) ||
		   (game.Player2 != nil && game.Player2.ID == playerID) {
			playerGames = append(playerGames, game)
		}
	}
	return playerGames
}

// GetGameByID возвращает игру по ID
func (gs *GameService) GetGameByID(games []*Game, gameID string) *Game {
	for _, game := range games {
		if game.ID == gameID {
			return game
		}
	}
	return nil
}

// IsGameFinished проверяет, завершена ли игра
func (gs *GameService) IsGameFinished(game *Game) bool {
	return game.Status == GameStatusFinished
}

// GetWinner возвращает победителя игры
func (gs *GameService) GetWinner(game *Game) *Player {
	if game.Status != GameStatusFinished {
		return nil
	}
	return game.Winner
}

// IsDraw проверяет, является ли игра ничьей
func (gs *GameService) IsDraw(game *Game) bool {
	return game.Status == GameStatusFinished && game.Winner == nil
}

// GetCurrentTurn возвращает игрока, чей сейчас ход
func (gs *GameService) GetCurrentTurn(game *Game) *Player {
	return game.CurrentTurn
}

// ValidateGameState проверяет валидность состояния игры
func (gs *GameService) ValidateGameState(game *Game) error {
	if game.Player1 == nil {
		return errors.New("game must have at least one player")
	}
	
	if game.Board == nil {
		return errors.New("game must have a board")
	}
	
	if game.Status == GameStatusActive && game.Player2 == nil {
		return errors.New("active game must have two players")
	}
	
	if game.Status == GameStatusActive && game.CurrentTurn == nil {
		return errors.New("active game must have current turn")
	}
	
	return nil
}

// GetGameStatistics возвращает статистику игры
func (gs *GameService) GetGameStatistics(game *Game) GameStatistics {
	stats := GameStatistics{
		GameID:     game.ID,
		Status:     game.Status,
		TotalMoves: 0,
		Duration:   time.Duration(0),
	}
	
	// Подсчитываем количество ходов
	if game.Board != nil {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if game.Board.GetCell(i, j) != "" {
					stats.TotalMoves++
				}
			}
		}
	}
	
	// Вычисляем длительность игры
	if game.StartedAt != nil {
		endTime := time.Now()
		if game.FinishedAt != nil {
			endTime = *game.FinishedAt
		}
		stats.Duration = endTime.Sub(*game.StartedAt)
	}
	
	return stats
}

// GameStatistics - статистика игры
type GameStatistics struct {
	GameID     string        `json:"game_id"`
	Status     GameStatus    `json:"status"`
	TotalMoves int           `json:"total_moves"`
	Duration   time.Duration `json:"duration"`
}

// GetRandomAvailablePosition возвращает случайную доступную позицию
func (gs *GameService) GetRandomAvailablePosition(game *Game) int {
	if game.Board == nil {
		return -1
	}
	
	var available []int
	for i := 0; i < 9; i++ {
		if game.Board.IsValidPosition(i) {
			available = append(available, i)
		}
	}
	
	if len(available) == 0 {
		return -1
	}
	
	rand.Seed(time.Now().UnixNano())
	return available[rand.Intn(len(available))]
} 