-- +goose Up
CREATE TABLE
  IF NOT EXISTS "todos" (
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    "user_id" uuid NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    "title" text NOT NULL,
    "done" boolean NOT NULL DEFAULT false,
    "created_at" timestamp
    with
      time zone DEFAULT now () NOT NULL,
    "updated_at" timestamp
    with
      time zone DEFAULT now () NOT NULL
  );

-- +goose Down
DROP TABLE IF EXISTS "todos";
