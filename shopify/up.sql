CREATE TABLE IF NOT EXISTS tokens (
    shop_name VARCHAR(255) NOT NULL,
    account_id VARCHAR(255) NOT NULL,
    token VARCHAR(255) NOT NULL,
    CONSTRAINT tokens_unique UNIQUE (shop_name, account_id)
);
 

CREATE TABLE IF NOT EXISTS orders (
    id VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    shop_name VARCHAR(255) NOT NULL,
    account_id VARCHAR(255) NOT NULL,
    order_id VARCHAR(255) NOT NULL,
    total_price FLOAT NOT NULL
);

