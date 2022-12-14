CREATE TABLE users(
    id INTEGER PRIMARY KEY
);

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
    user_id INTEGER REFERENCES users(id),
    service_id INTEGER NOT NULL,
    is_positive BOOLEAN NOT NULL,
    price NUMERIC(18, 2),
    description VARCHAR(256),
    execution_date DATE NOT NULL DEFAULT CURRENT_DATE
);

INSERT INTO users(id) VALUES(0);
INSERT INTO users(id) VALUES(1);
INSERT INTO users(id) VALUES(2);