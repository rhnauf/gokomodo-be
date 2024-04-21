-- +goose Up
CREATE TABLE sellers (
  id SERIAL PRIMARY KEY,
  email VARCHAR(100) NOT NULL,
  name VARCHAR(50) NOT NULL,
  password TEXT,
  address_pickup VARCHAR(150) NOT NULL
);

-- +goose Down
DROP TABLE sellers;
