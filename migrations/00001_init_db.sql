-- +goose Up
CREATE TABLE verification (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  is_removed BOOLEAN NOT NULL DEFAULT FALSE
);

-- +goose Down
DROP TABLE verification;
