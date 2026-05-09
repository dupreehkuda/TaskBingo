CREATE INDEX IF NOT EXISTS games_user1_finished_idx ON games (user1_id, finished);
CREATE INDEX IF NOT EXISTS games_user2_finished_idx ON games (user2_id, finished);
