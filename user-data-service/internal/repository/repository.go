package repository

import (
	"context"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
	"go.uber.org/zap"
)

// Enumeration for friend status column
const (
	_ = iota
	FriendRequested
	FriendResponse
	FriendFriend
)

// repository provide a connection with database
type repository struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

// New creates a new instance of database layer and migrates it
func New(path string, logger *zap.Logger) *repository {
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

	return &repository{
		pool:   pool,
		logger: logger,
	}
}

// CreateSchema gets and executes needed schema
func (r repository) CreateSchema(path string) {
	schema, err := os.ReadFile(path)
	if err != nil {
		r.logger.Error("Error occurred while getting migration schema", zap.Error(err))
	}

	_, err = r.pool.Exec(context.Background(), string(schema))
	if err != nil {
		r.logger.Error("Error occurred while executing schema", zap.Error(err))
	}
}
