package app

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// MigrateUp applies all up migrations from path against dsn. ErrNoChange is non-fatal.
func MigrateUp(dsn, path string) error {
	m, err := migrate.New("file://"+path, dsn)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	src, dbErr := m.Close()
	if src != nil {
		return src
	}
	return dbErr
}
