-- +goose Up
-- SQL in this section is executed when the migration is applied.

-- +goose StatementBegin
-- pgcrypto installation https://www.chrisnewland.com/postgresql-9-pgcrypto-debian--370 
CREATE EXTENSION pgcrypto;

CREATE OR REPLACE function randomStringByLength(length integer) returns text as
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
    id          BIGINT         PRIMARY KEY  DEFAULT         nextval('users_id_seq'),
    name        VARCHAR(100)   NOT NULL,
    nickname    VARCHAR(50)    DEFAULT      '',
    username    VARCHAR(50)    NOT NULL     UNIQUE,
    gender      VARCHAR(10)    DEFAULT      'male',
    status      VARCHAR(20)    DEFAULT      'single',
    birthdate   TIMESTAMP      DEFAULT      NULL,
    blood_type  VARCHAR(5)     DEFAULT      NULL,

    email       VARCHAR(100)   NOT NULL     UNIQUE,
    password    VARCHAR(50)    DEFAULT      '',
    photo       VARCHAR(255)   DEFAULT      '',
    is_active   BOOLEAN        DEFAULT      FALSE,
    timezone    VARCHAR(50)    DEFAULT      'Asia/Jakarta'  NOT NULL,
    language    VARCHAR(50)    DEFAULT      'Indonesia'     NOT NULL,
    signature   VARCHAR(100)   DEFAULT      'Regards',
    api_token  VARCHAR(100)   DEFAULT      TO_CHAR(now(), 'DD-MON-YYYY_HH24-MI-SS-US')||'_'||randomStringByLength(50),

    deleted_at  TIMESTAMP      DEFAULT      NULL,
    created_at  TIMESTAMP      DEFAULT      CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP      DEFAULT      CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS users;
DROP SEQUENCE IF EXISTS users_id_seq;
DROP FUNCTION IF EXISTS randomStringByLength(integer);
DROP EXTENSION IF EXISTS pgcrypto;
