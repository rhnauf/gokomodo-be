-- +goose Up
CREATE TABLE orders (
  id SERIAL PRIMARY KEY,
  buyer_id BIGINT REFERENCES buyers(id) ON DELETE CASCADE ON UPDATE CASCADE,
  seller_id BIGINT REFERENCES sellers(id) ON DELETE CASCADE ON UPDATE CASCADE,
  address_source VARCHAR(150) NOT NULL,
  address_destination VARCHAR(150) NOT NULL,
  product_id BIGINT REFERENCES products(id) ON DELETE CASCADE ON UPDATE CASCADE,
  qty INT default 0,
  price DOUBLE PRECISION NOT NULL,
  total_price DOUBLE PRECISION NOT NULL,
  status VARCHAR(10) default 'pending'
);

-- +goose Down
DROP TABLE orders;
