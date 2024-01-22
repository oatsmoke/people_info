-- +goose Up
-- +goose StatementBegin
CREATE TABLE people
(
    id          serial primary key not null,
    Name        varchar            not null,
    Surname     varchar            not null,
    Patronymic  varchar,
    Age         int                not null,
    Gender      varchar            not null,
    Nationality varchar            not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE people;
-- +goose StatementEnd
