package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SslMode  string
}

func InitDB(dbconfig *DatabaseConfig) (*pgxpool.Pool, error) {
	conn_str := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", dbconfig.Host, dbconfig.Port, dbconfig.User, dbconfig.Password, dbconfig.DBName, dbconfig.SslMode)
	config, err := pgxpool.ParseConfig(conn_str)
	config.MaxConns=100
	
	if err != nil {
		return nil, err
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}
	return pool, nil

}
