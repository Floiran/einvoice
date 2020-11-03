package main

import (
	"log"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		log.Println("Creating table users")
		_, err := db.Exec(`
			CREATE TABLE users (
				id VARCHAR ( 100 ) PRIMARY KEY,
				name VARCHAR ( 100 ) NOT NULL,
				service_account_key TEXT,
				email VARCHAR ( 100 )
			);
		`)

		return err
	}, func(db migrations.DB) error {
		log.Println("Dropping table users...")
		_, err := db.Exec(`DROP TABLE users;`)

		return err
	})
}
