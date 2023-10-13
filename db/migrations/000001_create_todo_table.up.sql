CREATE TABLE IF NOT EXISTS todos
(
    id           uuid PRIMARY KEY,
    title        VARCHAR(100) NOT NULL,
    description  VARCHAR(500) NOT NULL,
    is_completed bool         NOT NULL,
    created_at   timestamptz  NOT NULL,
    created_by   varchar(100) NOT NULL,
    updated_at   timestamptz  NOT NULL,
    updated_by   varchar(100) NOT NULL,
    deleted_at   timestamptz  NULL,
    deleted_by   VARCHAR(100) NOT NULL DEFAULT ''
);