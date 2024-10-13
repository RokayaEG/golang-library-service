-- +goose Up
-- +goose StatementBegin
CREATE TABLE genres (
    id int NOT NULL UNIQUE AUTO_INCREMENT,
    code varchar(255) NOT NULL UNIQUE,
    name varchar(255) NOT NULL,
    slug varchar(255),
    description text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE genres;
-- +goose StatementEnd
