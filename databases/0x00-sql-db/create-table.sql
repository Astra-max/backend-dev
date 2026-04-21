CREATE TABLE IF NOT EXISTS users(
    user_id  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    create_at TIMESTAMPTZ DEFAULT NOW(),
);