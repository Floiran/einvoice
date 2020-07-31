CREATE TABLE invoices (
	id serial PRIMARY KEY,
	sender VARCHAR ( 50 ) NOT NULL,
	receiver VARCHAR ( 50 ) NOT NULL
);