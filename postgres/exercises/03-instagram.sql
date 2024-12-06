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