package main

import (
	"github.com/jackc/pgx"
)

type PostgresConfig struct {
	Host           string
	Port           uint16
	Database       string
	User           string
	Password       string
	RuntimeParams  map[string]string
	AfterConnect   func(*pgx.Conn) error
	MaxConnections int
}

func GenerateConfig(conf PostgresConfig) pgx.ConnPoolConfig {
	return pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:          conf.Host,
			Port:          conf.Port,
			Database:      conf.Database,
			User:          conf.User,
			Password:      conf.Password,
			RuntimeParams: conf.RuntimeParams,
		},
		MaxConnections: conf.MaxConnections,
		AfterConnect:   conf.AfterConnect,
	}
}

type PostgresConn struct {
	Conn   *pgx.ConnPool
	Config pgx.ConnPoolConfig
}

func NewPostgresConn(conf pgx.ConnPoolConfig) (*PostgresConn, error) {
	conn, err := pgx.NewConnPool(conf)
	if err != nil {
		return nil, err
	}

	pgconn := &PostgresConn{
		Conn:   conn,
		Config: conf,
	}

	return pgconn, nil
}

func (p *PostgresConn) PostgresVersion() {
	return
}

func (p *PostgresConn) ShowConfig() {
	return
}

func (p *PostgresConn) ShowConnections() {
	return
}
