-- +goose Up
-- +goose StatementBegin
create table if not exists questions (
    id serial primary key,
    q_text text not null,
    created_at timestamp not null default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists questions
-- +goose StatementEnd
