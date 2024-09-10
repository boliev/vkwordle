-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE nouns
(
    word  VARCHAR(15) NOT NULL,
    type    SMALLINT NOT NULL DEFAULT 5
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
    DROP TABLE nouns;
-- +goose StatementEnd
