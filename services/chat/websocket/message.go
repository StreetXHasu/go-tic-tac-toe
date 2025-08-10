package websocket

import (
	"encoding/json"
	"time"
)

// MessageType message type
type MessageType string

const (
	MessageTypeChat     MessageType = "chat"
	MessageTypeJoin     MessageType = "join"
	MessageTypeLeave    MessageType = "leave"
	MessageTypeGameMove MessageType = "game_move"
	MessageTypeSystem   MessageType = "system"
	MessageTypeError    MessageType = "error"
)

// Message message structure
type Message struct {
	Type      MessageType `json:"type"`
	Content   string      `json:"content"`
	Username  string      `json:"username"`
	GameID    string      `json:"game_id"`
	Timestamp time.Time   `json:"timestamp"`
	Data      interface{} `json:"data,omitempty"`
}

// NewMessage creates a new message
func NewMessage(msgType MessageType, content, username, gameID string) *Message {
	return &Message{
		Type:      msgType,
		Content:   content,
		Username:  username,
		GameID:    gameID,
		Timestamp: time.Now(),
	}
}

// ToJSON converts message to JSON
func (m *Message) ToJSON() ([]byte, error) {
	return json.Marshal(m)
}

// FromJSON creates message from JSON
func FromJSON(data []byte) (*Message, error) {
	var msg Message
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}

// ChatMessage structure for chat messages
type ChatMessage struct {
	Message
	IsPrivate bool `json:"is_private"`
}

// NewChatMessage creates a new chat message
func NewChatMessage(content, username, gameID string, isPrivate bool) *ChatMessage {
	return &ChatMessage{
		Message:   *NewMessage(MessageTypeChat, content, username, gameID),
		IsPrivate: isPrivate,
	}
}

// GameMoveMessage structure for game move messages
type GameMoveMessage struct {
	Message
	Position int    `json:"position"`
	Symbol   string `json:"symbol"`
}

// NewGameMoveMessage creates a new game move message
func NewGameMoveMessage(position int, symbol, username, gameID string) *GameMoveMessage {
	return &GameMoveMessage{
		Message:  *NewMessage(MessageTypeGameMove, "", username, gameID),
		Position: position,
		Symbol:   symbol,
	}
}

// SystemMessage structure for system messages
type SystemMessage struct {
	Message
	Action string `json:"action"`
}

// NewSystemMessage creates a new system message
func NewSystemMessage(action, content, gameID string) *SystemMessage {
	return &SystemMessage{
		Message: *NewMessage(MessageTypeSystem, content, "System", gameID),
		Action:  action,
	}
}

// ErrorMessage structure for error messages
type ErrorMessage struct {
	Message
	ErrorCode string `json:"error_code"`
}

// NewErrorMessage creates a new error message
func NewErrorMessage(errorCode, content, gameID string) *ErrorMessage {
	return &ErrorMessage{
		Message:   *NewMessage(MessageTypeError, content, "System", gameID),
		ErrorCode: errorCode,
	}
}

// JoinMessage structure for join messages
type JoinMessage struct {
	Message
	PlayerCount int `json:"player_count"`
}

// NewJoinMessage creates a new join message
func NewJoinMessage(username, gameID string, playerCount int) *JoinMessage {
	return &JoinMessage{
		Message:     *NewMessage(MessageTypeJoin, username+" joined the game", username, gameID),
		PlayerCount: playerCount,
	}
}

// LeaveMessage structure for leave messages
type LeaveMessage struct {
	Message
	PlayerCount int `json:"player_count"`
}

// NewLeaveMessage creates a new leave message
func NewLeaveMessage(username, gameID string, playerCount int) *LeaveMessage {
	return &LeaveMessage{
		Message:     *NewMessage(MessageTypeLeave, username+" left the game", username, gameID),
		PlayerCount: playerCount,
	}
} 