CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    balance_id SERIAL REFERENCES balances(id),
    account_id SERIAL REFERENCES accounts(id),
    order_id SERIAL REFERENCES orders(id)
);