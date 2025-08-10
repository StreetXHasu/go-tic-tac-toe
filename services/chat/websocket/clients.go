package websocket

import (
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Client represents a WebSocket client
type Client struct {
	ID       string
	Username string
	GameID   string
	Conn     *websocket.Conn
	Hub      *Hub
	Send     chan []byte
	mu       sync.Mutex
}

// NewClient creates a new client
func NewClient(id, username, gameID string, conn *websocket.Conn, hub *Hub) *Client {
	return &Client{
		ID:       id,
		Username: username,
		GameID:   gameID,
		Conn:     conn,
		Hub:      hub,
		Send:     make(chan []byte, 256),
	}
}

// ReadPump handles incoming messages from client
func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(pongWait)
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(pongWait)
		return nil
	})

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		// Process message
		c.Hub.HandleMessage(c, message)
	}
}

// WritePump sends messages to client
func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(writeWait)
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add additional messages to queue
			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write([]byte{'\n'})
				w.Write(<-c.Send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(writeWait)
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// SendMessage sends a message to client
func (c *Client) SendMessage(message []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	select {
	case c.Send <- message:
	default:
		close(c.Send)
		c.Hub.Unregister <- c
	}
}

// GetGameID returns the client's game ID
func (c *Client) GetGameID() string {
	return c.GameID
}

// GetUsername returns the client's username
func (c *Client) GetUsername() string {
	return c.Username
}

// IsConnected checks if client is connected
func (c *Client) IsConnected() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	select {
	case <-c.Send:
		return false
	default:
		return true
	}
}

// Close closes the client connection
func (c *Client) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	close(c.Send)
	c.Conn.Close()
} 