-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS workouts (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  -- user_id (add a foreign key later if needed)
  title TEXT NOT NULL,
  description TEXT,
  duration_minutes INTEGER NOT NULL,
  calories_burned INTEGER,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS workouts;
-- +goose StatementEnd
