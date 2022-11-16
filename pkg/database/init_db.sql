CREATE DATABASE balance_manager;

CREATE TABLE balances(
    id SERIAL PRIMARY KEY,
    sum NUMERIC(18, 2)
);

CREATE TABLE accounts(
    id SERIAL PRIMARY KEY,
    sum NUMERIC(18, 2)
);

CREATE TABLE orders(
    id SERIAL PRIMARY KEY,
    service_id INTEGER NOT NULL,
    price NUMERIC(18, 2),
    description VARCHAR(256),
    execution_date DATE NOT NULL DEFAULT CURRENT_DATE
);

CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    balance_id SERIAL UNIQUE REFERENCES balances(id),
    account_id SERIAL REFERENCES accounts(id),
    order_id SERIAL REFERENCES orders(id)
);