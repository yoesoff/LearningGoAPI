-- +goose Up
-- SQL in this section is executed when the migration is applied.

-- +goose StatementBegin
CREATE SEQUENCE companies_id_seq  START 1;
CREATE TABLE companies(
    id          BIGINT          PRIMARY KEY     DEFAULT nextval('companies_id_seq'),
    name        VARCHAR(255)    NOT NULL,
    photo       VARCHAR(255)    DEFAULT '',
    tagline     VARCHAR(255)    DEFAULT '',
    currency    VARCHAR(20)     DEFAULT 'USD',
    
    deleted_at  TIMESTAMP       DEFAULT NULL,   
    created_at  TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP       DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS companies;
DROP SEQUENCE IF EXISTS companies_id_seq;
