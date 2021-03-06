-- +goose Up
-- SQL in this section is executed when the migration is applied.

-- +goose StatementBegin
CREATE SEQUENCE employees_id_seq  START 1;
CREATE TABLE employees(
    id          BIGINT          PRIMARY KEY     DEFAULT nextval('employees_id_seq'),
    name        VARCHAR(255)    NOT NULL,
    photo       VARCHAR(255)    DEFAULT '',
    
    deleted_at  TIMESTAMP       DEFAULT NULL,   
    created_at  TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP       DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS employees;
DROP SEQUENCE IF EXISTS employees_id_seq;
