package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Conn *pgxpool.Pool
}

func New(ctx context.Context, connString string) (*Postgres, error) {
	conn, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return &Postgres{Conn: conn}, nil
}

func (p *Postgres) Close(ctx context.Context) error {
	p.Conn.Close()
	return nil
}
