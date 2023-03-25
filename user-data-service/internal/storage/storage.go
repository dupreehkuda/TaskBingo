package storage

import (
	"context"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
	"go.uber.org/zap"

	"github.com/dupreehkuda/TaskBingo/user-data-service/internal/models"
)

// Stored is interface for storage
type Stored interface {
	GetUserData(userID string) (*models.GetUserDataResponse, error)
	CheckDuplicateUser(username, email string) (bool, error)
	CreateUser(userID, username, email, passwordHash, passwordSalt, city string) error
	LoginUser(username string) (*models.LoginUserResponse, error)

	GetRatedPacks() ([]string, error)
	LikePack(userID, pack string, inc int) error
	RatePack(userID, pack string, inc int) error
	AssignNewPack(userID, packID string, packName string) error

	GetAllUsers() (*[]models.AllUsers, error)
	AcceptFriend(userID, friendID string) error
	DeleteFriend(userID, friendID string) error
	RequestFriend(userID, friendID string) error
	
	CreateGame(game *models.Game) error
}

// Enumeration for friend status column
const (
	_ = iota
	Requested
	Response
	Friend
)

// storage provide a connection with database
type storage struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

// New creates a new instance of database layer and migrates it
func New(path string, logger *zap.Logger) *storage {
	// Wait until database initialize in container
	time.Sleep(2 * time.Second)

	config, err := pgxpool.ParseConfig(path)
	if err != nil {
		logger.Error("Unable to parse config", zap.Error(err))
	}

	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxUUID.Register(conn.TypeMap())
		return nil
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		logger.Error("Unable to connect to database", zap.Error(err))
	}

	return &storage{
		pool:   pool,
		logger: logger,
	}
}

// CreateSchema gets and executes needed schema
func (s storage) CreateSchema() {
	schema, err := os.ReadFile("./2023-03-04-migrate.sql")
	if err != nil {
		s.logger.Error("Error occurred while getting migration schema", zap.Error(err))
	}

	_, err = s.pool.Exec(context.Background(), string(schema))
	if err != nil {
		s.logger.Error("Error occurred while executing schema", zap.Error(err))
	}
}
