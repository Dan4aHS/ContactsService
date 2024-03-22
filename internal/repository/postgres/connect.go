package postgres

import (
	"ContactsService/internal/config"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

const driverName = "pgx"

func ConnectPostgres(cfg *config.Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Database,
	)
	connCfg, err := pgx.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}
	db, err := sql.Open(driverName, stdlib.RegisterConnConfig(connCfg))
	if err != nil {
		return nil, err
	}
	dbx := sqlx.NewDb(db, driverName)
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return dbx, nil
}
