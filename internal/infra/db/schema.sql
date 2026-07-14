CREATE TABLE users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    active boolean NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);