-- data depends on sql/04-instagram-data.sql
-- to import the data, refer to 03-build-instagram-app.md

-- 1. Selet the 3 users with the highest IDs from the users table
SELECT * FROM users
ORDER BY id DESC
LIMIT 3;

-- 2. Join the users and posts table. 
-- Show the username of user ID 200 and the captions of all posts they have created.
SELECT username, caption
FROM posts
JOIN users ON posts.user_id = users.id
WHERE posts.user_id = 200;

-- 3. Show each username and the number of likes that they have created.
SELECT username, COUNT(*) AS "Number of Likes"
FROM likes
JOIN users ON likes.user_id = users.id
GROUP BY users.username;

-- 4. Show the username of users who were tagged in a caption or photo before Jan 7th, 2010
-- Also show the date they were tagged.
SELECT username, created_at
FROM users
JOIN (
    SELECT user_id, created_at FROM caption_tags
    UNION ALL
    SELECT user_id, created_at FROM photo_tags
) AS tags ON tags.user_id = users.id
WHERE users.created_at < '2010-01-07';

-- 5. Suggest a list of followers from the user with id 1000's friends. 