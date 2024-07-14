-- +goose Up
-- +goose StatementBegin
CREATE TABLE games
(
    id      SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    puzzle  VARCHAR(15) NOT NULL,
    type    SMALLINT NOT NULL DEFAULT 5,
    status  SMALLINT NOT NULL DEFAULT 0,
    words   VARCHAR(15)[] NOT NULL DEFAULT '{}'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE games;
-- +goose StatementEnd
