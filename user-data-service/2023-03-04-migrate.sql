CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY NOT NULL UNIQUE,
    username text NOT NULL UNIQUE,
    email text NOT NULL UNIQUE,
    city text NOT NULL,
    wins integer DEFAULT 0,
    lose integer DEFAULT 0,
    bingo integer DEFAULT 0,
    likedpacks uuid[] DEFAULT '{}',
    ratedpacks uuid[] DEFAULT '{}',
    registered timestamptz NOT NULL,
    games uuid[] DEFAULT '{}'
);

CREATE TABLE IF NOT EXISTS login (
    id uuid PRIMARY KEY NOT NULL UNIQUE,
    passwordhash text NOT NULL UNIQUE,
    passwordsalt text NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS packs (
    id uuid PRIMARY KEY NOT NULL UNIQUE,
    title text NOT NULL,
    tasks text[] NOT NULL,
    rating integer NOT NULL DEFAULT 0,
    liked integer NOT NULL DEFAULT 1,
    played integer NOT NULL DEFAULT 0,
    creator uuid NOT NULL,
    created timestamptz NOT NULL,
    last_played timestamptz
);

CREATE TABLE IF NOT EXISTS friends (
    id uuid NOT NULL,
    friend_id uuid NOT NULL,
    status integer NOT NULL,
    wins integer NOT NULL DEFAULT 0,
    loses integer NOT NULL DEFAULT 0,
    since timestamptz
);

CREATE TABLE IF NOT EXISTS games (
    id uuid PRIMARY KEY NOT NULL UNIQUE,
    user1_id uuid NOT NULL,
    user2_id uuid NOT NULL,
    pack_id uuid NOT NULL,
    status integer NOT NULL,
    user1_bingo integer default 0,
    user2_bingo integer default 0,
    winner uuid,
    numbers integer[] default '{}',
    user1_numbers integer[] default '{}',
    user2_numbers integer[] default '{}',
    created timestamptz,
    accepted timestamptz,
    finished timestamptz
);

ALTER TABLE "login" ADD FOREIGN KEY ("id") REFERENCES "users" ("id");
ALTER TABLE "ratings" ADD FOREIGN KEY ("creator") REFERENCES "users" ("id");
ALTER TABLE "friends" ADD FOREIGN KEY ("id") REFERENCES "users" ("id");
ALTER TABLE "packs" ADD FOREIGN KEY ("creator") REFERENCES "users" ("id");