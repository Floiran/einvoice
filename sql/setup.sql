DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'format') THEN
        CREATE TYPE format AS ENUM ('ubl2.1', 'd16b');
    END IF;
END$$;

CREATE TABLE IF NOT EXISTS invoices (
	id SERIAL PRIMARY KEY,
	sender VARCHAR ( 100 ) NOT NULL,
	receiver VARCHAR ( 100 ) NOT NULL,
	format format NOT NULL,
	price DECIMAL NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS attachments (
	id SERIAL PRIMARY KEY,
	name VARCHAR ( 100 ) NOT NULL,
	invoice_id integer NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	CONSTRAINT fk_invoice
      FOREIGN KEY(invoice_id)
	  REFERENCES invoices(id)
);
