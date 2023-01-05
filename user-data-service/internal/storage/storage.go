package storage

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

// storage provide a connection with database
type storage struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

var schema = `
CREATE TABLE IF NOT EXISTS "users" (
   "id" text PRIMARY KEY NOT NULL UNIQUE,
   "login" text NOT NULL UNIQUE,
   "email" text NOT NULL UNIQUE,
   "city" text NOT NULL,
   "registered" timestamptz NOT NULL ,
   "wins" integer DEFAULT 0,
   "lose" integer DEFAULT 0,
   "scoreboard" integer DEFAULT 0,
   "friends" text[],
   "packs" text[]
);

CREATE TABLE IF NOT EXISTS "login" (
	"id" text PRIMARY KEY NOT NULL UNIQUE,
	"passwordhash" text NOT NULL UNIQUE,
	"passwordsalt" text NOT NULL UNIQUE 
);

CREATE TABLE IF NOT EXISTS "ratings" (
	"id" text PRIMARY KEY NOT NULL UNIQUE,
	"rating" integer NOT NULL DEFAULT 0,
	"liked" integer NOT NULL DEFAULT 1,
	"played" integer NOT NULL DEFAULT 0,
	"creator" text NOT NULL,
	"created" timestamptz NOT NULL
);

ALTER TABLE "login" ADD FOREIGN KEY ("id") REFERENCES "users" ("id");
ALTER TABLE "ratings" ADD FOREIGN KEY ("creator") REFERENCES "users" ("id")`

// New creates a new instance of database layer and migrates it
func New(path string, logger *zap.Logger) *storage {
	// Wait until database initialize in container
	time.Sleep(2 * time.Second)

	config, err := pgxpool.ParseConfig(path)
	if err != nil {
		logger.Error("Unable to parse config", zap.Error(err))
	}

	conn, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		logger.Error("Unable to connect to database", zap.Error(err))
	}

	return &storage{
		pool:   conn,
		logger: logger,
	}
}

// CreateSchema executes needed schema
func (s storage) CreateSchema() {
	_, err := s.pool.Exec(context.Background(), schema)
	if err != nil {
		s.logger.Error("Error occurred while executing schema", zap.Error(err))
	}

	s.logger.Info("Launched with pgx. Database created.")
}
