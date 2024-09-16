-- +goose Up
-- +goose StatementBegin
CREATE TABLE options (
    id SERIAL PRIMARY KEY,
    poll_id INT NOT NULL,
    body VARCHAR(200) NOT NULL,
    vote_count INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_poll_id FOREIGN KEY (poll_id) REFERENCES polls(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE options;
-- +goose StatementEnd
