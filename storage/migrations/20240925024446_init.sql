-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id UUID NOT NULL PRIMARY KEY UNIQUE,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL
);

CREATE TABLE tasks (
    id SERIAL NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    level_id INT NOT NULL
);

CREATE TABLE users_groups (
    user_id UUID REFERENCES users (id) ON DELETE CASCADE NOT NULL,
    group_id SERIAL REFERENCES groups (id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE tasks_groups (
    task_id SERIAL REFERENCES tasks (id) ON DELETE CASCADE NOT NULL,
    group_id SERIAL REFERENCES groups (id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE groups (
    id SERIAL NOT NULL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE level_task (
    level_id SERIAL REFERENCES level_progress (id) ON DELETE CASCADE NOT NULL,
    task_id SERIAL REFERENCES tasks (id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE level_progress (
    id SERIAL PRIMARY KEY NOT NULL,
    name TEXT NOT NULL
);

CREATE TABLE refresh_tokens (
    id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE NOT NULL ,
    refresh_token_hash TEXT UNIQUE NOT NULL,
    ip VARCHAR(45) NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL ,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE refresh_tokens;
DROP TABLE users_groups;
DROP TABLE tasks_groups;
DROP TABLE level_task;
DROP TABLE users;
DROP TABLE groups;
DROP TABLE tasks;
DROP TABLE level_progress;
-- +goose StatementEnd
