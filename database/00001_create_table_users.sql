-- +goose Up
-- SQL in this section is executed when the migration is applied.

-- +goose StatementBegin
-- pgcrypto installation https://www.chrisnewland.com/postgresql-9-pgcrypto-debian--370 
CREATE EXTENSION pgcrypto;

Create or replace function random_string_by_length(length integer) returns text as
$$
DECLARE
  characters TEXT := 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
  bytes BYTEA := gen_random_bytes(length);
  l INT := length(characters);
  i INT := 0;
  output TEXT := '';
BEGIN
  WHILE i < length LOOP
    output := output || substr(characters, get_byte(bytes, i) % l + 1, 1);
    i := i + 1;
  END LOOP;
  RETURN output;
END;
$$ LANGUAGE plpgsql VOLATILE;


CREATE SEQUENCE users_id_seq  START 1;
CREATE TABLE users(
    id          BIGINT         PRIMARY KEY     DEFAULT nextval('users_id_seq'::regclass),
    name        VARCHAR(255)   NOT NULL,
    username    VARCHAR(255)   NOT NULL UNIQUE,
    email       VARCHAR(255)   NOT NULL UNIQUE,
    password    VARCHAR(255)   DEFAULT  '',
    is_active   BOOLEAN        DEFAULT  FALSE,
    timezone    VARCHAR(100)   DEFAULT  'Asia/Jakarta' NOT NULL,
    language    VARCHAR(100)   DEFAULT  'Indonesia'    NOT NULL,
    signature   VARCHAR(100)   DEFAULT  NULL,
    
    deleted_at  TIMESTAMP      DEFAULT  NULL,
    created_at  TIMESTAMP      DEFAULT  CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP      DEFAULT  CURRENT_TIMESTAMP,

    rest_token  VARCHAR(100)   DEFAULT  TO_CHAR(now(), 'DD-MON-YYYY_HH24-MI-SS-US')||'_'||random_string_by_length(50)
);
-- +goose StatementEnd

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS users;
DROP SEQUENCE IF EXISTS users_id_seq;
DROP FUNCTION IF EXISTS random_string_by_length(integer);
DROP EXTENSION IF EXISTS pgcrypto;
