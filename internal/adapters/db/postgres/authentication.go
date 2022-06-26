package postgres

import (
	"context"
	"github.com/jackc/pgx/v4"
	"ms-hw/internal/core"
	"ms-hw/internal/core/aggregate"
)

type authenticationStorage struct {
	db *pgx.Conn
}

func NewPostgresRepo(db *pgx.Conn) *authenticationStorage {
	return &authenticationStorage{db: db}
}

func (s *authenticationStorage) GetCredentials(ctx context.Context, person *aggregate.Person) (res *core.Credentials, err error) {
	_, err = s.db.Exec(ctx, "")
	if err != nil {
		return nil, err
	}
	return res, nil
}
