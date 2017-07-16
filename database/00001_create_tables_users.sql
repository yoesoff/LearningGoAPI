-- +goose Up
-- SQL in this section is executed when the migration is applied.

-- +goose StatementBegin
CREATE SEQUENCE users_id_seq  START 1;
CREATE TABLE users(
    id          BIGINT      PRIMARY KEY     DEFAULT nextval('users_id_seq'::regclass),
    name        CHAR(255)                   NOT NULL,
    username    CHAR(255)                   NOT NULL UNIQUE,
    email       CHAR(255)                   NOT NULL UNIQUE,
    password    CHAR(255)                   NOT NULL,
    is_active   BOOLEAN     DEFAULT         FALSE,
    timezone    CHAR(100)   DEFAULT         'Asia/Jakarta' NOT NULL,
    language    CHAR(100)   DEFAULT         'Indonesia'    NOT NULL,
    signature   TEXT        DEFAULT         'Kind regards' NOT NULL,
    deleted_at  TIMESTAMP,   
    created_at  TIMESTAMP   DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP   DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS users;
DROP SEQUENCE IF EXISTS users_id_seq;
