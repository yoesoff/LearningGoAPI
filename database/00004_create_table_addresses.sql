-- +goose Up
-- SQL in this section is executed when the migration is applied.

-- +goose StatementBegin
CREATE SEQUENCE addresses_id_seq  START 1;
CREATE TABLE addresses(
    id          BIGINT          PRIMARY KEY     DEFAULT nextval('addresses_id_seq'),
    name        VARCHAR(255)    NOT NULL,
    photo       VARCHAR(255)    DEFAULT '',
    street      VARCHAR(255)    DEFAULT '',
    city        VARCHAR(255)    DEFAULT '',
    state       VARCHAR(255)    DEFAULT '',
    zip         VARCHAR(50)     DEFAULT '',
    country     VARCHAR(50)     DEFAULT '',

    deleted_at  TIMESTAMP       DEFAULT NULL,   
    created_at  TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP       DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS addresses;
DROP SEQUENCE IF EXISTS addresses_id_seq;
