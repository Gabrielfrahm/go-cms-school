DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'cms-school') THEN
        CREATE DATABASE "cms-school";
    END IF;
END
$$;

DROP TABLE IF EXISTS user_tokens;

CREATE TABLE user_tokens(
    "user_id" uuid REFERENCES users(id) ON DELETE CASCADE,
    "token" TEXT NOT NULL,
    "refresh_token"  TEXT NOT NULL,
    "expires_at" TIMESTAMP(3),
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP
);
