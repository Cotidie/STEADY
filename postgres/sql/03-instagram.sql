CREATE TABLE "users" (
  "id"          SERIAL PRIMARY KEY,
  "username"    VARCHAR(30),
  "bio"         TEXT,
  "avatar"      TEXT,
  "phone"       TEXT,
  "email"       TEXT,
  "password"    TEXT,
  "status"      TEXT,
  "created_at"  TIMESTAMP DEFAULT (now()),
  "updated_at"  TIMESTAMP DEFAULT (now())
);

CREATE TABLE "posts" (
  "id"      SERIAL PRIMARY KEY,
  "url"     VARCHAR(200) NOT NULL,
  -- if a user is deleted, so is a post.
  "user_id" INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  "caption" VARCHAR(240),
  "lat"     REAL NOT NULL CHECK(lat >= -90 AND lat <= 90),
  "lng"     REAL NOT NULL CHECK(lng >= -180 AND lng <= 180)
);

CREATE TABLE "comments" (
  "id"          SERIAL PRIMARY KEY,
  "contents"    TEXT NOT NULL,
  "user_id"     INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  "post_id"     INTEGER NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
  "created_at"  TIMESTAMP DEFAULT (now()),
  "updated_at"  TIMESTAMP DEFAULT (now())
);

CREATE TABLE "likes" (
  "id"          SERIAL PRIMARY KEY,
  "user_id"     INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  -- polymorphic association over posts and comments -> can be NULL
  "post_id"     INTEGER REFERENCES posts(id) ON DELETE CASCADE,
  "comment_id"  INTEGER REFERENCES comments(id) ON DELETE CASCADE,
  CHECK(
	COALESCE((post_id)::BOOLEAN::INTEGER, 0)
	+
	COALESCE((comment_id)::BOOLEAN::INTEGER, 0)
	= 1
  ),
  UNIQUE(user_id, post_id, comment_id)
);

CREATE TABLE "caption_tags" (
  "id"      SERIAL PRIMARY KEY,
  "post_id" INTEGER NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
  "user_id" INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  UNIQUE(post_id, user_id)
);

CREATE TABLE "photo_tags" (
  "id"      SERIAL PRIMARY KEY,
  "x"       REAL NOT NULL DEFAULT 0,
  "y"       REAL NOT NULL DEFAULT 0,
  "post_id" INTEGER NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
  "user_id" INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  UNIQUE(post_id, user_id)
);

CREATE TABLE "hashtags" (
  "id"          SERIAL PRIMARY KEY,
  "title"       TEXT NOT NULL UNIQUE,
  "created_at"  TIMESTAMP DEFAULT (now())
);

CREATE TABLE "post_hashtags" (
  "id"          SERIAL PRIMARY KEY,
  "post_id"     INTEGER NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
  "hashtag_id"  INTEGER NOT NULL REFERENCES hashtags(id) ON DELETE CASCADE,
  UNIQUE(post_id, hashtag_id)
);

CREATE TABLE "followers" (
  "id"          SERIAL PRIMARY KEY,
  "leader_id"   INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  "follower_id" INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  UNIQUE(leader_id, follower_id)
);