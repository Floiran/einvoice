package main

import (
	"log"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		log.Println("Creating table invoices")
		_, err := db.Exec(`
			CREATE TYPE format AS ENUM ('ubl2.1', 'd16b');
			CREATE TABLE invoices (
				id SERIAL PRIMARY KEY,
				created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
				sender VARCHAR ( 100 ) NOT NULL,
				receiver VARCHAR ( 100 ) NOT NULL,
				format format NOT NULL,
				price DECIMAL NOT NULL
			);
		`)
		if err != nil {
			return err
		}

		log.Println("Creating table attachments")
		_, err = db.Exec(`
			CREATE TABLE attachments (
				id SERIAL PRIMARY KEY,
				name VARCHAR ( 100 ) NOT NULL,
				invoice_id integer NOT NULL,
				created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
				CONSTRAINT fk_invoice
				  FOREIGN KEY(invoice_id)
				  REFERENCES invoices(id)
			);
		`)

		return err
	}, func(db migrations.DB) error {
		log.Println("Dropping table attachments...")
		_, err := db.Exec(`DROP TABLE attachments;`)

		if err != nil {
			return err
		}

		log.Println("Dropping table invoices...")
		_, err = db.Exec(`
			DROP TABLE invoices;
			DROP TYPE format;
		`)

		return err
	})
}
