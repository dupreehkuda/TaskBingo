package storage

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

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

var schema = `
CREATE TABLE IF NOT EXISTS users (
   id text PRIMARY KEY NOT NULL UNIQUE,
   login text NOT NULL UNIQUE,
   email text NOT NULL UNIQUE,
   city text NOT NULL,
   wins integer DEFAULT 0,
   lose integer DEFAULT 0,
   bingo integer DEFAULT 0,
   likedpacks text[] DEFAULT '{}',
   ratedpacks text[] DEFAULT '{}',
   registered timestamptz NOT NULL,
   activegames uuid[] DEFAULT '{}',
   pastgames uuid[] DEFAULT '{}',
   requestedgames uuid[] DEFAULT '{}',
   createdgames uuid[] DEFAULT '{}'
);

CREATE TABLE IF NOT EXISTS login (
	id text PRIMARY KEY NOT NULL UNIQUE,
	passwordhash text NOT NULL UNIQUE,
	passwordsalt text NOT NULL UNIQUE 
);

CREATE TABLE IF NOT EXISTS ratings (
	id text PRIMARY KEY NOT NULL UNIQUE,
	rating integer NOT NULL DEFAULT 0,
	liked integer NOT NULL DEFAULT 1,
	played integer NOT NULL DEFAULT 0,
	creator text NOT NULL,
	created timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS friends (
	id text NOT NULL,
	friend text NOT NULL,
	status integer NOT NULL,
	wins integer NOT NULL DEFAULT 0,
	loses integer NOT NULL DEFAULT 0,
	since timestamptz
);

CREATE TABLE IF NOT EXISTS games (
  	uuid uuid NOT NULL,
  	user1_id text NOT NULL,
  	user2_id text NOT NULL,
  	pack_id text NOT NULL,
  	status integer NOT NULL,
  	user1_bingo integer default 0,
  	user2_bingo integer default 0,
  	winner text,
  	numbers integer[] default '{}',
  	user1_numbers integer[] default '{}',
  	user2_numbers integer[] default '{}',
  	created timestamptz
);

ALTER TABLE "login" ADD FOREIGN KEY ("id") REFERENCES "users" ("id");
ALTER TABLE "ratings" ADD FOREIGN KEY ("creator") REFERENCES "users" ("id");
ALTER TABLE "friends" ADD FOREIGN KEY ("id") REFERENCES "users" ("id");`

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
}
