CREATE TABLE orders(
    id SERIAL PRIMARY KEY,
    user_id SERIAL REFERENCES orders(id),
    service_id INTEGER NOT NULL,
    price NUMERIC(18, 2),
    description VARCHAR(256),
    execution_date DATE NOT NULL DEFAULT CURRENT_DATE
);