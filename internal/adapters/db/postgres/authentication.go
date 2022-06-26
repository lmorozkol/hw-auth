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

func (s *authenticationStorage) GetCredentials(ctx context.Context, person *aggregate.Person) (res []*core.Credentials, err error) {
	rows, err := s.db.Query(ctx, "select * from credentials where login=$1 and password=$2",
		person.Login,
		person.Password,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var row core.Credentials
		if err = rows.Scan(&row.ID, &row.Timestamp, &row.Login, &row.Password); err != nil {
			return nil, err
		}

		res = append(res, &row)
	}

	return res, nil
}
