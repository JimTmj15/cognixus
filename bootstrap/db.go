package bootstrap

import (
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewDb(env *Env) *sqlx.DB {
	conn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", env.DBUser, env.DBPass, env.DBHost, env.DBName)
	db, err := sqlx.Open("pgx", conn)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)
	return db
}
