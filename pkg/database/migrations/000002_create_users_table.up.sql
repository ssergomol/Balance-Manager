CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    balance_id SERIAL UNIQUE REFERENCES balances(id)
);