CREATE TABLE balances(
    id SERIAL PRIMARY KEY,
    user_id INTEGER UNIQUE REFERENCES users(id),
    sum NUMERIC(18, 2)
);