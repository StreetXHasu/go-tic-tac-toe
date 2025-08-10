package websocket

import (
	"encoding/json"
	"log"
	"sync"
	"time"
)

const (
	// Time to wait for message write
	writeWait = 10 * time.Second

	// Time to wait for pong message
	pongWait = 60 * time.Second

	// Period for sending ping messages
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size
	maxMessageSize = 512
)

// Hub manages WebSocket connections
type Hub struct {
	// Registered clients by games
	clients map[string]map[*Client]bool

	// Channels for client registration/unregistration
	Register   chan *Client
	Unregister chan *Client

	// Channel for sending messages
	Broadcast chan *Message

	// Mutex for safe access to clients
	mu sync.RWMutex
}

// NewHub creates a new hub
func NewHub() *Hub {
	return &Hub{
		clients:    make(map[string]map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message),
	}
}

// Run starts the hub
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.registerClient(client)

		case client := <-h.Unregister:
			h.unregisterClient(client)

		case message := <-h.Broadcast:
			h.broadcastMessage(message)
		}
	}
}

// registerClient registers a new client
func (h *Hub) registerClient(client *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	gameID := client.GetGameID()
	if h.clients[gameID] == nil {
		h.clients[gameID] = make(map[*Client]bool)
	}
	h.clients[gameID][client] = true

	// Send join message
	joinMsg := NewJoinMessage(client.GetUsername(), gameID, len(h.clients[gameID]))
	h.broadcastToGame(gameID, joinMsg)

	log.Printf("Client %s joined game %s", client.GetUsername(), gameID)
}

// unregisterClient unregisters a client
func (h *Hub) unregisterClient(client *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	gameID := client.GetGameID()
	if h.clients[gameID] != nil {
		if _, ok := h.clients[gameID][client]; ok {
			delete(h.clients[gameID], client)
			close(client.Send)

			// Send leave message
			leaveMsg := NewLeaveMessage(client.GetUsername(), gameID, len(h.clients[gameID]))
			h.broadcastToGame(gameID, leaveMsg)

			// Remove game if no clients
			if len(h.clients[gameID]) == 0 {
				delete(h.clients, gameID)
			}

			log.Printf("Client %s left game %s", client.GetUsername(), gameID)
		}
	}
}

// broadcastMessage sends message to all clients in game
func (h *Hub) broadcastMessage(message *Message) {
	h.broadcastToGame(message.GameID, message)
}

// broadcastToGame sends message to all clients in specific game
func (h *Hub) broadcastToGame(gameID string, message interface{}) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	clients, exists := h.clients[gameID]
	if !exists {
		return
	}

	// Convert message to JSON
	jsonData, err := json.Marshal(message)
	if err != nil {
		log.Printf("Error marshaling message: %v", err)
		return
	}

	// Send message to all clients in game
	for client := range clients {
		select {
		case client.Send <- jsonData:
		default:
			close(client.Send)
			delete(clients, client)
		}
	}
}

// HandleMessage processes incoming messages from client
func (h *Hub) HandleMessage(client *Client, message []byte) {
	var msg Message
	if err := json.Unmarshal(message, &msg); err != nil {
		// Send error message
		errorMsg := NewErrorMessage("INVALID_JSON", "Invalid message format", client.GetGameID())
		if jsonData, err := errorMsg.ToJSON(); err == nil {
			client.SendMessage(jsonData)
		}
		return
	}

	// Set missing fields
	msg.Username = client.GetUsername()
	msg.GameID = client.GetGameID()
	msg.Timestamp = time.Now()

	// Process message based on type
	switch msg.Type {
	case MessageTypeChat:
		h.handleChatMessage(client, &msg)
	case MessageTypeGameMove:
		h.handleGameMoveMessage(client, &msg)
	case MessageTypeJoin:
		h.handleJoinMessage(client, &msg)
	case MessageTypeLeave:
		h.handleLeaveMessage(client, &msg)
	default:
		// Send error message
		errorMsg := NewErrorMessage("UNKNOWN_TYPE", "Unknown message type", client.GetGameID())
		if jsonData, err := errorMsg.ToJSON(); err == nil {
			client.SendMessage(jsonData)
		}
	}
}

// handleChatMessage processes chat messages
func (h *Hub) handleChatMessage(client *Client, msg *Message) {
	// Check message length
	if len(msg.Content) > 200 {
		errorMsg := NewErrorMessage("MESSAGE_TOO_LONG", "Message too long", client.GetGameID())
		if jsonData, err := errorMsg.ToJSON(); err == nil {
			client.SendMessage(jsonData)
		}
		return
	}

	// Create chat message
	chatMsg := NewChatMessage(msg.Content, client.GetUsername(), client.GetGameID(), false)
	
	// Send message to all clients in game
	h.broadcastToGame(client.GetGameID(), chatMsg)
}

// handleGameMoveMessage processes game move messages
func (h *Hub) handleGameMoveMessage(client *Client, msg *Message) {
	// Here you can add additional logic for processing moves
	gameMoveMsg := NewGameMoveMessage(0, "X", client.GetUsername(), client.GetGameID())
	h.broadcastToGame(client.GetGameID(), gameMoveMsg)
}

// handleJoinMessage processes join messages
func (h *Hub) handleJoinMessage(client *Client, msg *Message) {
	h.mu.RLock()
	playerCount := len(h.clients[client.GetGameID()])
	h.mu.RUnlock()

	joinMsg := NewJoinMessage(client.GetUsername(), client.GetGameID(), playerCount)
	h.broadcastToGame(client.GetGameID(), joinMsg)
}

// handleLeaveMessage processes leave messages
func (h *Hub) handleLeaveMessage(client *Client, msg *Message) {
	h.mu.RLock()
	playerCount := len(h.clients[client.GetGameID()])
	h.mu.RUnlock()

	leaveMsg := NewLeaveMessage(client.GetUsername(), client.GetGameID(), playerCount)
	h.broadcastToGame(client.GetGameID(), leaveMsg)
}

// GetClientCount returns number of clients in game
func (h *Hub) GetClientCount(gameID string) int {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if clients, exists := h.clients[gameID]; exists {
		return len(clients)
	}
	return 0
}

// GetActiveGames returns list of active games
func (h *Hub) GetActiveGames() []string {
	h.mu.RLock()
	defer h.mu.RUnlock()

	var games []string
	for gameID := range h.clients {
		games = append(games, gameID)
	}
	return games
} 