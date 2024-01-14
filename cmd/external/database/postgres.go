package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"simple-ecommerce/cmd/internal/config"
	"time"
)

func ConnectPostgres(cfg config.DBConfig) (db *sqlx.DB, err error) {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
	)

	db, err = sqlx.Open("postgres", dsn)
	if err != nil {
		return
	}

	// TEST PING DB
	if err = db.Ping(); err != nil {
		return
	}

	// Connection Pool
	db.SetConnMaxIdleTime(time.Duration(cfg.ConnectionPool.MaxIdleConnection) * time.Second)
	db.SetConnMaxLifetime(time.Duration(cfg.ConnectionPool.MaxLifetimeConnection) * time.Second)
	db.SetMaxOpenConns(int(cfg.ConnectionPool.MaxOpenConnetcion))
	db.SetMaxIdleConns(int(cfg.ConnectionPool.MaxIdleConnection))

	return
}
