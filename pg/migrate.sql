CREATE TABLE IF NOT EXISTS orders (
  id BIGINT PRIMARY KEY,
  title TEXT,
  customer_id INTEGER NOT NULL,
  price INTEGER NOT NULL,
  created_at timestamp without time zone NOT NULL DEFAULT NOW(),
  updated_at timestamp without time zone
);
