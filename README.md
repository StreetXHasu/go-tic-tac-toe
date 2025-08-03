# Tic-Tac-Toe Microservices

Advanced Tic-Tac-Toe game with customizable board size and win conditions. Play solo or with friends through invitation links, real-time chat, and leaderboards.

## Game Features

- **Customizable Board Size**: Play on boards from 3x3 to 10x10
- **Flexible Win Conditions**: Set custom win sequence length (3-5 in a row)
- **Solo Mode**: Practice against AI with adjustable difficulty
- **Multiplayer**: Play with friends through invitation links
- **Real-time Chat**: Communicate with opponents during games
- **Leaderboards**: Track wins, losses, and rankings
- **Game History**: Review past games and moves
- **Responsive Design**: Works on desktop and mobile devices

## Architecture

- **Frontend**: HTML + Nginx (simplified version for MVP)
- **Gateway**: Fiber (Go) - API Gateway
- **Services**:
  - Auth Service - authentication and authorization
  - User Service - user management and profiles
  - Game Service - game logic with DDD approach
  - Chat Service - real-time player communication (WebSocket)
- **Database**: PostgreSQL with separate databases for each service
- **Shared Package**: common code in pkg/ to reduce duplication
- **Testing**: BDD tests with Godog for each service

## Ports

- Frontend: `http://localhost:3000`
- Gateway: `http://localhost:8080`
- Auth Service: `http://localhost:8081`
- User Service: `http://localhost:8082`
- Game Service: `http://localhost:8083`
- Chat Service: `http://localhost:8084`
- PostgreSQL: `localhost:5432`

## 🚀 Getting Started

### Quick Start

```bash
# Start all services
cd infra && docker-compose up -d

# Check status
docker ps

# View logs
cd infra && docker-compose logs -f

# Stop
cd infra && docker-compose down
```

### Health Check

After startup, verify service availability:

```bash
# Frontend
curl http://localhost:3000

# API Gateway
curl http://localhost:8080

# Health checks
curl http://localhost:8080/health
curl http://localhost:8081/health
curl http://localhost:8082/health
curl http://localhost:8083/health
curl http://localhost:8084/health
```

### Web Interface

Open in browser: **http://localhost:3000**

You will see a beautiful HTML page with status check for all services and game interface.

### Game Setup

1. **Create Account**: Register to track your statistics
2. **Choose Game Mode**: 
   - Solo vs AI (adjustable difficulty)
   - Multiplayer (invite friends)
3. **Configure Board**: Select size (3x3 to 10x10) and win sequence length
4. **Start Playing**: Make moves and chat with opponents
5. **View Leaderboards**: Check rankings and achievements

## Testing

```bash
# Run BDD tests for specific service
cd services/auth && godog run tests/

# Run all tests
make test

# Run tests with coverage
make test-coverage
```

## Logging

All services log to shared volume `logs`. Logs are saved to files:

- `frontend.log` - frontend logs (Nginx)
- `gateway.log` - API Gateway logs
- `auth-service.log` - authentication service logs
- `user-service.log` - user service logs
- `game-service.log` - game service logs
- `chat-service.log` - chat service logs
- `postgres.log` - database logs

### Log management commands:

```bash
# View logs for specific service
make logs-gateway
make logs-auth
make logs-game

# View all logs
make logs-all

# View only errors
make logs-error

# Log statistics
make logs-stats

# Clear logs
make logs-clear
```

## Databases

Created automatically:
- `auth_db` - for authentication service
- `user_db` - for user service
- `game_db` - for game service
- `chat_db` - for chat service

## Project Structure

```
go-tic-tac-toe/
├── README.md                   # Project documentation
├── Makefile                    # Build, test and management commands
├── infra/                      # Infrastructure and Docker
│   ├── docker-compose.yaml     # All services configuration
│   ├── scripts/
│   │   ├── init-multiple-databases.sh  # Database initialization
│   │   └── view-logs.sh                # Log management
│   └── README.md               # Infrastructure documentation
├── pkg/                        # Common code for all services
│   ├── go.mod                  # Go module
│   ├── config/
│   │   └── config.go           # Common configuration
│   ├── middleware/
│   │   ├── cors.go             # CORS middleware
│   │   ├── auth.go             # Authentication
│   │   └── logging.go          # Request logging
│   ├── database/
│   │   └── postgres.go         # PostgreSQL connection
│   ├── utils/
│   │   ├── jwt.go              # JWT tokens
│   │   ├── response.go         # HTTP responses
│   │   └── logger.go           # Structured logging
│   └── models/
│       └── common.go           # Common data models
├── frontend/                   # Frontend (HTML + Nginx)
│   ├── Dockerfile              # Image build
│   ├── nginx.conf              # Nginx configuration
│   └── index.html              # Main page
├── gateway/                    # API Gateway (Fiber)
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   ├── main.go                 # Entry point
│   ├── handlers/
│   │   └── proxy.go            # Request proxying
│   ├── routes/
│   │   └── routes.go           # API routes
│   └── tests/                  # Tests
│       ├── gateway.feature     # BDD scenarios
│       └── gateway_test.go     # Unit tests
└── services/                   # Microservices
    ├── auth/                   # Authentication service
    │   ├── Dockerfile
    │   ├── go.mod
    │   ├── go.sum
    │   ├── main.go
    │   ├── handlers/
    │   │   ├── auth.go         # Authentication handlers
    │   │   └── health.go       # Health check
    │   ├── models/
    │   │   └── user.go         # User models
    │   ├── routes/
    │   │   └── routes.go       # API routes
    │   └── tests/              # BDD tests
    │       ├── auth.feature
    │       └── auth_test.go
    ├── user/                   # User service
    │   ├── Dockerfile
    │   ├── go.mod
    │   ├── go.sum
    │   ├── main.go
    │   ├── handlers/
    │   │   ├── user.go         # User handlers
    │   │   └── health.go
    │   ├── models/
    │   │   └── user.go
    │   ├── routes/
    │   │   └── routes.go
    │   └── tests/
    │       ├── user.feature
    │       └── user_test.go
    ├── game/                   # Game service (DDD approach)
    │   ├── Dockerfile
    │   ├── go.mod
    │   ├── go.sum
    │   ├── main.go
    │   ├── domain/             # Domain logic
    │   │   ├── game.go         # Game aggregate
    │   │   ├── board.go        # Board value object
    │   │   ├── player.go       # Player entity
    │   │   ├── move.go         # Move value object
    │   │   └── logic.go        # Game logic
    │   ├── handlers/
    │   │   ├── game.go         # HTTP handlers
    │   │   └── health.go
    │   ├── models/
    │   │   └── game.go         # Database models
    │   ├── routes/
    │   │   └── routes.go
    │   └── tests/
    │       ├── game.feature
    │       └── game_test.go
    └── chat/                   # Chat service (WebSocket)
        ├── Dockerfile
        ├── go.mod
        ├── go.sum
        ├── main.go
        ├── handlers/
        │   ├── chat.go         # HTTP handlers
        │   └── health.go
        ├── models/
        │   └── message.go      # Message models
        ├── routes/
        │   └── routes.go
        ├── websocket/          # WebSocket components
        │   ├── hub.go          # Connection management
        │   ├── clients.go      # Client management
        │   └── message.go      # Message types
        └── tests/
            ├── chat.feature
            └── chat_test.go
```

## Architecture Benefits

### Shared Package (pkg/)
- **config/** - common configuration for all services
- **middleware/** - reusable middleware (CORS, authentication, logging)
- **database/** - common functions for PostgreSQL
- **utils/** - utilities (JWT, HTTP responses, logging)
- **models/** - common data models

### DDD Approach for Game Service
- **domain/** - domain logic isolated from infrastructure
- **Aggregates** - Game as main aggregate with board state and players
- **Value Objects** - Board, Move, WinCondition as immutable objects
- **Entities** - Player as entity with identifier and statistics
- **Domain Services** - game logic, AI opponent, win detection in logic.go

### BDD Testing
- **tests/** - test folder for each service
- **.feature files** - scenarios in Gherkin language
- **Godog** - BDD framework for Go
- **Test coverage** - all critical scenarios

### Game Modes and Features
- **Classic Mode**: Traditional 3x3 board with 3-in-a-row win
- **Extended Mode**: Larger boards (4x4 to 10x10) with custom win sequences
- **AI Opponent**: Multiple difficulty levels for solo play
- **Invitation System**: Share game links with friends
- **Real-time Updates**: Live game state synchronization
- **Chat Integration**: In-game messaging between players
- **Statistics Tracking**: Win/loss ratios and performance metrics

### Using Shared Package
```go
// In each service
import (
    "github.com/your-org/go-tic-tac-toe/pkg/config"
    "github.com/your-org/go-tic-tac-toe/pkg/middleware"
    "github.com/your-org/go-tic-tac-toe/pkg/database"
    "github.com/your-org/go-tic-tac-toe/pkg/utils"
)
```



## Commands

### Makefile Commands

```bash
# Build and testing
make build              # Build all Go services
make test               # Run unit tests for all services
make test-bdd           # Run BDD tests for services
make test-coverage      # Run tests with coverage
make clean              # Clean build artifacts
make lint               # Run linters

# Dependencies installation
make deps               # Install Go dependencies
make install-godog      # Install Godog for BDD tests
make install-lint       # Install linters
make setup              # Install all tools and dependencies

# Docker commands
make docker-build       # Build Docker images
make docker-up          # Start all services
make docker-down        # Stop all services
make docker-logs        # View Docker logs

# Log management commands
make logs-frontend      # Frontend logs
make logs-gateway       # API Gateway logs
make logs-auth          # Auth service logs
make logs-user          # User service logs
make logs-game          # Game service logs
make logs-chat          # Chat service logs
make logs-postgres      # Database logs
make logs-all           # All logs in real-time
make logs-error         # Only errors
make logs-info          # Only info messages
make logs-debug         # Only debug messages
make logs-stats         # Log statistics
make logs-clear         # Clear all logs

# Help
make help               # Show all available commands
```

### Docker Commands

```bash
# Basic commands
docker-compose up -d    # Start all services in background
docker-compose down     # Stop all services
docker-compose logs -f  # View logs in real-time
docker-compose build    # Rebuild images

# Service management
docker-compose restart gateway    # Restart gateway
docker-compose restart frontend   # Restart frontend
docker-compose logs gateway       # Logs only for gateway

# Cleanup
docker-compose down -v            # Stop with volume removal
docker system prune -a            # Clean all unused resources
```

### Health Check

```bash
# Check container status
docker ps

# Check API endpoints
curl http://localhost:8080        # Gateway
curl http://localhost:8080/health # Health check
curl http://localhost:3000        # Frontend

# Check individual services
curl http://localhost:8081/health # Auth service
curl http://localhost:8082/health # User service
curl http://localhost:8083/health # Game service
curl http://localhost:8084/health # Chat service
``` 