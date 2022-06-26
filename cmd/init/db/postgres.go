package db

import (
	"context"
	"ms-hw/internal/config"

	"github.com/jackc/pgx/v4"
)

func Conn(ctx context.Context, cfg *config.Cfg) (*pgx.Conn, error) {
	dbUrl := "postgres://" + cfg.DbLogin + ":" + cfg.DbPass + "@" + cfg.DbAddress + ":" + cfg.DbPort + "/" + cfg.DbName
	conn, err := pgx.Connect(ctx, dbUrl)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
