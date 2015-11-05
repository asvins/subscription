CREATE TABLE subscriptions (
	email varchar(30) CONSTRAINT subscription_pk PRIMARY KEY,
	cpf TEXT,
	address TEXT,
	delivery_address TEXT,
	credit_card TEXT,
	phone TEXT
);

CREATE TABLE subscribers (
	email varchar(30) CONSTRAINT subscriber_pk PRIMARY KEY,
	payment_status INTEGER,
	last_payed TIMESTAMP WITHOUT TIME ZONE,
	next_payment TIMESTAMP WITHOUT TIME ZONE
);
