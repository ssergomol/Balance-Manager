CREATE TABLE orders(
    id SERIAL PRIMARY KEY,
    service_id INTEGER NOT NULL,
    price NUMERIC(18, 2),
    description TEXT(256),
    execution_date DATE NOT NULL DEFAULT CURRENT_DATE
);