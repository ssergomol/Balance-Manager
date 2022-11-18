CREATE TABLE accounts(
    id INTEGER PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    sum NUMERIC(18, 2)
);