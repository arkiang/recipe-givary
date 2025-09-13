-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS recipes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    making_time TEXT NOT NULL,
    serves TEXT NOT NULL,
    ingredients TEXT NOT NULL,
    cost INTEGER NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO recipes (title, making_time, serves, ingredients, cost, created_at, updated_at) VALUES
    ('Chicken Curry', '45 min', '4 people', 'onion, chicken, seasoning', 1000, '2016-01-10 12:10:12', '2016-01-10 12:10:12');

INSERT INTO recipes (title, making_time, serves, ingredients, cost, created_at, updated_at) VALUES
    ('Rice Omelette', '30 min', '2 people', 'onion, egg, seasoning, soy sauce', 700, '2016-01-11 13:10:12', '2016-01-11 13:10:12');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS recipes;
-- +goose StatementEnd
