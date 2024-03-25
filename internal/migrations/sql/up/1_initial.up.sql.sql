DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'cms-school') THEN
        CREATE DATABASE "cms-school";
    END IF;
END
$$;

-- Create Table teste
DROP TABLE IF EXISTS todo;
CREATE TABLE todo(
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    "title" VARCHAR(50) NOT NULL,
    "description" VARCHAR(50) NOT NULL,
    "done" BOOLEAN NOT NULL DEFAULT false
);