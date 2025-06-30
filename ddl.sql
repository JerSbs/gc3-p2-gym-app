-- Users Table
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email VARCHAR(100) NOT NULL UNIQUE,
  full_name VARCHAR(100) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  weight INTEGER NOT NULL,
  height INTEGER NOT NULL
);

-- Workouts Table
CREATE TABLE workouts (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  description TEXT NOT NULL,
  user_id INTEGER NOT NULL,
  CONSTRAINT fk_workout_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Exercises Table
CREATE TABLE exercises (
  id SERIAL PRIMARY KEY,
  workout_id INTEGER NOT NULL,
  name VARCHAR(100) NOT NULL,
  description VARCHAR(255) NOT NULL,
  CONSTRAINT fk_exercise_workout FOREIGN KEY (workout_id) REFERENCES workouts(id) ON DELETE CASCADE
);

-- Exercise Logs Table
CREATE TABLE exercise_logs (
  id SERIAL PRIMARY KEY,
  exercise_id INTEGER NOT NULL,
  user_id INTEGER NOT NULL,
  set_count INTEGER NOT NULL,
  rep_count INTEGER NOT NULL,
  weight INTEGER NOT NULL,
  CONSTRAINT fk_log_exercise FOREIGN KEY (exercise_id) REFERENCES exercises(id) ON DELETE CASCADE,
  CONSTRAINT fk_log_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
