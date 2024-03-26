DO $$
BEGIN
    IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'cms-school') THEN
        CREATE DATABASE "cms-school";
    END IF;
END
$$;


DROP TABLE IF EXISTS profile_permissions;
DROP TABLE IF EXISTS user_permissions;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS profiles;


CREATE TABLE profiles (
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    "name" TEXT,
    "type_user" TEXT NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP(3)
);

CREATE TABLE users(
    "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    "name" TEXT,
    "email" TEXT UNIQUE,
    "password" TEXT,
    "type_user" TEXT NOT NULL,
    "profile_id" uuid REFERENCES profiles(id),
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP(3)
);

CREATE TABLE user_permissions (
    "user_id" uuid REFERENCES users(id),
    "users" INT NOT NULL, 
    "classes" INT NOT NULL,
    "profiles" INT NOT NULL,
    "lessons" INT NOT NULL,
    PRIMARY KEY (user_id)
);


CREATE TABLE profile_permissions (
    "profile_id" uuid REFERENCES profiles(id),
    "users" INT NOT NULL, 
    "classes" INT NOT NULL,
    "profiles" INT NOT NULL,
    "lessons" INT NOT NULL,
    PRIMARY KEY (profile_id)
);