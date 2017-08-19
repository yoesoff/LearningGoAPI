-- +goose Up
-- SQL in this section is executed when the migration is applied.

-- +goose StatementBegin
CREATE SEQUENCE contacts_id_seq  START 1;
CREATE TABLE contacts(
    id          BIGINT          PRIMARY KEY     DEFAULT nextval('contacts_id_seq'),
    name            VARCHAR(255)   NOT NULL,
    website         VARCHAR(255)   DEFAULT '',
    email           VARCHAR(255)   DEFAULT '',
    fax             VARCHAR(30)    DEFAULT '',
    land_phone      VARCHAR(30)    DEFAULT '',
    mobile_phone    VARCHAR(30)    DEFAULT '',
    notes           VARCHAR(255)   DEFAULT '',

    deleted_at  TIMESTAMP       DEFAULT NULL,   
    created_at  TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP       DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS contacts;
DROP SEQUENCE IF EXISTS contacts_id_seq;
