-- +goose Up

CREATE TABLE feeds(
	id UUID PRIMARY KEY,
	created_at TIMESTAMP DEFAULT LOCALTIMESTAMP,
	updated_at TIMESTAMP,
	name TEXT NOT NULL,
	url TEXT UNIQUE NOT NULL,
	user_id UUID,
	CONSTRAINT fk_users
	FOREIGN KEY (user_id)
	REFERENCES users(id)  ON DELETE CASCADE
);

-- +goose Down

DROP TABLE feeds;