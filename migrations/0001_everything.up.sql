CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS transactions (
  id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  price decimal,
  status text NOT NULL,
  create_time TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS events (
  id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  event_type text,
  payload bytea,
  create_time TIMESTAMPTZ NOT NULL DEFAULT now()  
);
