-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS workout_entries (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  workout_id INTEGER NOT NULL,
  exercise_name TEXT NOT NULL,
  sets INTEGER NOT NULL,
  reps INTEGER,
  duration_seconds INTEGER,
  weight REAL,
  notes TEXT,
  order_index INTEGER,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_workout FOREIGN KEY (workout_id) REFERENCES workouts(id) ON DELETE CASCADE,
  CONSTRAINT valid_workout_entry CHECK (
      (reps IS NOT NULL OR duration_seconds IS NOT NULL) AND
      (reps IS NULL OR duration_seconds IS NULL)
  )
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS workout_entries;
-- +goose StatementEnd
