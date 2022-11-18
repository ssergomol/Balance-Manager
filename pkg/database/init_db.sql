CREATE DATABASE balance_manager;

CREATE TABLE balances(
    id SERIAL PRIMARY KEY,
    user_id INTEGER UNIQUE REFERENCES users(id),
    sum NUMERIC(18, 2)
);

CREATE TABLE accounts(
    id INTEGER PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    sum NUMERIC(18, 2)
);

CREATE TABLE orders(
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES orders(id),
    service_id INTEGER NOT NULL,
    price NUMERIC(18, 2),
    description VARCHAR(256),
    execution_date DATE NOT NULL DEFAULT CURRENT_DATE
);

CREATE TABLE users(
    id INTEGER PRIMARY KEY
);