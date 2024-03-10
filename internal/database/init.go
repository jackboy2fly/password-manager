package database

import (
	"fmt"
	"log"
	"time"

	"github.com/jackboy2fly/password-manager/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE IF NOT EXISTS service (
	id serial primary key,
	name varchar(50),
	secret varchar(150),
	created_at timestamp,
	updated_at timestamp
)`

type Service struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Secret    string    `db:"secret"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func Init() {
	connStr := fmt.Sprintf(
		"user=%s dbname=%s password=%s sslmode=%s",
		config.Config("DB_USERNAME"),
		config.Config("DB_NAME"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_SSL_MODE"),
	)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(schema)

	q := `INSERT INTO service
	(name, secret, created_at, updated_at)
	VALUES ($1, $2, $3, $4)`

	tx := db.MustBegin()
	tx.MustExec(q, config.Config("DB_FIRST_NAME"), config.Config("DB_FIRST_SECRET"), time.Now().UTC(), time.Now().UTC())
}
