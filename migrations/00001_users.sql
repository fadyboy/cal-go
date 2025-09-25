-- +goose Up
-- +goose StatementBegin
CREATE TABLEusers(
	id SERIAL PRIMARY KEY,
	email TEXT UNIQUE NOT NULL,
	password_hash TEXT NOT NULL
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLEusers;
-- +goose StatementEnd
