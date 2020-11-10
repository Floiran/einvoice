package main

import (
	"github.com/go-pg/migrations/v8"
	log "github.com/sirupsen/logrus"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		log.Println("Adding user columns")
		_, err := db.Exec(`
			ALTER TABLE users
            ADD COLUMN tax_id BIGINT,
            ADD COLUMN vat_number BIGINT;
		`)

		return err
	}, func(db migrations.DB) error {
		log.Println("Dropping users columns...")
		_, err := db.Exec(`
			ALTER TABLE users
			DROP COLUMN tax_id,
			DROP COLUMN vat_number;
		`)

		return err
	})
}
