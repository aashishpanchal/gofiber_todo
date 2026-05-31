-- +goose Up
CREATE TABLE
  IF NOT EXISTS "users" (
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    "name" varchar(255) NOT NULL,
    "email" varchar(255) NOT NULL UNIQUE,
    "password" text NOT NULL,
    "created_at" timestamp
    with
      time zone DEFAULT now () NOT NULL,
    "updated_at" timestamp
    with
      time zone DEFAULT now () NOT NULL
  );

-- +goose Down
DROP TABLE IF EXISTS "users";