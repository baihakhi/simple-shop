CREATE TABLE IF NOT EXISTS locations (
    location_id VARCHAR(50) NOT NULL PRIMARY KEY,
    address VARCHAR(255) NOT NULL,
    district VARCHAR(100) NOT NULL,
    city VARCHAR(100) NOT NULL,
    province VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
    user_id BIGSERIAL NOT NULL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    address VARCHAR(50),
    role VARCHAR(50) NOT NULL,
    balance INT DEFAULT 999000000,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    UNIQUE(username),
    FOREIGN KEY (address) REFERENCES locations(location_id)
      ON DELETE SET NULL ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS products (
    product_id BIGSERIAL NOT NULL PRIMARY KEY,
    code VARCHAR(50) NOT NULL,
    title VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    weight SMALLINT NOT NULL DEFAULT 1,
    stock SMALLINT NOT NULL DEFAULT 1,
    address VARCHAR(50),
    category VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    UNIQUE(code),
    FOREIGN KEY (address) REFERENCES locations(location_id) ON DELETE SET NULL ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS carts (
    cart_id BIGSERIAL NOT NULL PRIMARY KEY,
    user_id INT NOT NULL,
    product_id INT NOT NULL,
    amount SMALLINT NOT NULL DEFAULT 1,
    amount_price INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(product_id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS transactions (
    transaction_id INT NOT NULL PRIMARY KEY,
    cart_id INT NOT NULL,
    shipment_price INT NOT NULL DEFAULT 15000,
    origin VARCHAR(50) NOT NULL,
    destination VARCHAR(50) NOT NULL,
    code VARCHAR(50) NOT NULL,
    total_charge INT NOT NULL,
    status VARCHAR(50) DEFAULT 'CREATED',

    FOREIGN KEY (origin) REFERENCES locations(location_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (destination) REFERENCES locations(location_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (cart_id) REFERENCES carts(cart_id)
);

-- create function for automated updated_at on update records
CREATE  FUNCTION update_timestamp_func()
RETURNS TRIGGER
LANGUAGE plpgsql AS
'
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
';

DO $$
DECLARE
    t text;
BEGIN
    FOR t IN
        SELECT table_name FROM information_schema.columns WHERE column_name = 'updated_at'
    LOOP
        EXECUTE format('CREATE TRIGGER trigger_update_timestamp
                    BEFORE UPDATE ON %I
                    FOR EACH ROW EXECUTE PROCEDURE update_timestamp_func()', t,t);
    END loop;
END;
$$ language 'plpgsql';

-- adding locations record to the table
INSERT INTO locations (location_id, address, district, city, province) 
VALUES (
    'SLM001', 
    'Jl. Palagan Tentara Pelajar KM.8',
    'SLEMAN', 
    'SLEMAN',
    'YOGYAKARTA');

-- adding users record to the table with role ADMIN
INSERT INTO users (
    username,
    full_name, 
    address, 
    role, 
    password) 
VALUES (
    'simple', 
    'admin simple shop', 
    'SLM001',
    'ADMIN', 
    '$2a$12$T8I3jBhkxjJmZYF31/isBuCOTmrSq3jIJu8.R7ZcbhenPWCxHkP3G'); -- password: admin

-- adding products record to the table
INSERT INTO products (
    code,
    title, 
    price, 
    weight, 
    stock,
    address,
    category) 
VALUES (
    'S2000', 
    'Samsung S20+ Sein 8/128 Cosmic Black seken original - Garis banyak', 
    2999000,
    200, 
    99,
    'SLM001',
    'ELECTRONICS');
