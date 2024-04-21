-- +goose Up
CREATE TABLE products (
  id SERIAL PRIMARY KEY,
  product_name VARCHAR(100) NOT NULL,
  description TEXT,
  price DOUBLE PRECISION NOT NULL,
  seller_id BIGINT REFERENCES sellers(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- +goose Down
DROP TABLE products;
