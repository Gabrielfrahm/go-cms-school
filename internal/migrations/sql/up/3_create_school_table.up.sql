DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'cms-school') THEN
        CREATE DATABASE "cms-school";
    END IF;
END
$$;


DROP TABLE IF EXISTS addresses;

CREATE TABLE addresses(
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    "zip_code" TEXT NOT NULL,
    "city" TEXT NOT NULL,
    "address" TEXT NOT NULL,
    "number" TEXT NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP(3)
);

DROP TABLE IF EXISTS schools;

CREATE TABLE schools(
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    "name" TEXT NOT NULL,
    "address_id" uuid REFERENCES addresses(id) ON DELETE CASCADE,
    "director_id" uuid REFERENCES users(id) ON DELETE SET NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP(3)
);
