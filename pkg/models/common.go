package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

type User struct {
	BaseModel
	Email     string `json:"email" gorm:"uniqueIndex;not null"`
	Username  string `json:"username" gorm:"uniqueIndex;not null"`
	Password  string `json:"-" gorm:"not null"` // "-" means don't include in JSON
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsActive  bool   `json:"is_active" gorm:"default:true"`
}

type Game struct {
	BaseModel
	Player1ID   uint      `json:"player1_id" gorm:"not null"`
	Player2ID   *uint     `json:"player2_id"`
	Status      string    `json:"status" gorm:"default:'waiting'"` // waiting, active, finished
	WinnerID    *uint     `json:"winner_id"`
	Board       string    `json:"board" gorm:"default:'---------'"` // 9 characters: X, O, -
	CurrentTurn uint      `json:"current_turn" gorm:"default:1"`
	StartedAt   *time.Time `json:"started_at"`
	FinishedAt  *time.Time `json:"finished_at"`
}

type Message struct {
	BaseModel
	GameID  uint   `json:"game_id" gorm:"not null"`
	UserID  uint   `json:"user_id" gorm:"not null"`
	Content string `json:"content" gorm:"not null"`
} 