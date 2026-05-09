ALTER TABLE packs ADD COLUMN IF NOT EXISTS is_private boolean NOT NULL DEFAULT false;

ALTER TABLE games ALTER COLUMN user2_id DROP NOT NULL;
ALTER TABLE games ADD COLUMN IF NOT EXISTS kind text NOT NULL DEFAULT 'duo';

ALTER TABLE users ADD COLUMN IF NOT EXISTS solo_bingo integer NOT NULL DEFAULT 0;

CREATE TABLE IF NOT EXISTS game_comments (
    id uuid PRIMARY KEY NOT NULL,
    game_id uuid NOT NULL REFERENCES games(id) ON DELETE CASCADE,
    user_id uuid NOT NULL REFERENCES users(id),
    body text NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz
);

CREATE INDEX IF NOT EXISTS game_comments_game_user_idx ON game_comments(game_id, user_id);
