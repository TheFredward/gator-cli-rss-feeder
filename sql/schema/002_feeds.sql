-- +goose Up
CREATE TABLE feeds (
    id uuid PRIMARY KEY,
	created_at timestamp NOT NULL,
	updated_at timestamp NOT NULL,
	name varchar(255) UNIQUE,
	url varchar(255) UNIQUE,
	user_id uuid REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;
