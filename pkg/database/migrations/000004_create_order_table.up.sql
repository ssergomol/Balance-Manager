CREATE TABLE orders(
    id SERIAL PRIMARY KEY,
    service_id INTEGER,
    price NUMERIC(18, 2),
    description TEXT(256)
);