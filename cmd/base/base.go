package base

import (
	"context"
	"fmt"
	"os"
	"repo/cmd/base/items"
	"repo/cmd/base/users"
	"repo/cmd/models"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	db *pgxpool.Pool
}

type Memory struct {
	*Postgres
	*items.ItemsBase
	*users.UsersBase
}

func New(ctx context.Context) *Memory {
	return &Memory{
		NewPG(ctx),
		items.New(),
		users.New(),
	}
}

var (
	pgInstance *Postgres
	pgOnce     sync.Once
)

const connString = "postgres://postgres:test1@localhost:5432/repo"

func NewPG(ctx context.Context) *Postgres {
	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, connString)
		if err != nil {
			os.Exit(1)
		}
		pgInstance = &Postgres{db}
	})
	return pgInstance

}

func (p *Postgres) NewUser(ctx context.Context, user models.CreateUser) error {
	query := "INSERT INTO users(account, name, pass, email,  groupuser) VALUES ($1, $2, $3, $4, $5);"
	tag, err := p.db.Exec(ctx, query, user.Account, user.Name, user.Pass, user.Email, 0)
	fmt.Println(tag)
	fmt.Println(err)
	return err
}
