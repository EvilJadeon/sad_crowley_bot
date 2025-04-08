CREATE TABLE metrics (
  id BIGSERIAL NOT NULL PRIMARY KEY,
  user_id BIGINT NOT NULL REFERENCES users(id),
  week_id BIGINT NOT NULL REFERENCES weeks(id),
  completed_tasks INT NOT NULL,
  uncompleted_tasks INT NOT NULL,
  planned_tasks INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);