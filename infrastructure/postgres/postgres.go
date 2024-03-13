package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

type DB struct {
	Conn *pgx.Conn
}

type Config struct {
	User   string
	Pass   string
	Host   string
	Port   string
	DBName string
}

func NewDB(c Config) (*DB, error) {
	db := &DB{}
	if err := db.connect(c); err != nil {
		return nil, err
	}
	return db, nil
}

func (db *DB) connect(config Config) error {
	conn, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		config.User,
		config.Pass,
		config.Host,
		config.Port,
		config.DBName))
	if err != nil {
		return err
	}

	db.Conn = conn
	return nil
}

func (db *DB) Close() {
	defer func() {
		_ = db.Conn.Close(context.Background())
	}()
}
