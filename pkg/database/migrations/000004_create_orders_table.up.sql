CREATE TABLE orders(
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES orders(id),
    service_id INTEGER NOT NULL,
    is_positive BOOLEAN NOT NULL,
    price NUMERIC(18, 2),
    execution_date DATE NOT NULL DEFAULT CURRENT_DATE
);