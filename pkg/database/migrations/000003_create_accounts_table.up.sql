CREATE TABLE accounts(
    id SERIAL PRIMARY KEY,
    user_id SERIAL REFERENCES users(id),
    sum NUMERIC(18, 2)
);