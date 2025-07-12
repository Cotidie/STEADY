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
WITH RECURSIVE suggestions(leader_id, follower_id, depth) AS (
        SELECT leader_id, follower_id, 1
        FROM followers
        WHERE follower_id = 1000
    UNION
        SELECT followers.leader_id, followers.follower_id, depth+1
        FROM followers
        JOIN suggestions ON suggestions.leader_id = followers.follower_id
        WHERE depth < 3
)

SELECT id, username FROM users
JOIN suggestions ON suggestions.leader_id = users.id; 

-- 6. Show users by an descending order that the most tagged user would
-- come as the first row. A user can be tagged either by captions or photos.
SELECT username, COUNT(*)
FROM users
JOIN (
        SELECT user_id FROM photo_tags
    UNION ALL
        SELECt user_id FROM caption_tags
) AS tags ON users.id = tags.user_id
GROUP BY username
ORDER BY COUNT(*) DESC;

-- 7. Create a view that merges all the tags from captions and photos
-- and make a new column representing the type of its tag as: 'photo_tag' and 'caption_tag'
CREATE VIEW tags AS (
        SELECT id, created_at, user_id, post_id, 'photo_tag' AS type
        FROM photo_tags
    UNION ALL
        SELECT id, created_at, user_id, post_id, 'caption_tag' AS type 
        FROM caption_tags
);

-- 8. For each week, show the number of likes that posts and comments received.
-- Use the post and comment created_at date, not when the like was received.
CREATE MATERIALIZED VIEW weekly_likes AS (
    SELECT 
        DATE_TRUNC('week', COALESCE(posts.created_at, comments.created_at)) AS week ,
        COUNT(posts.id) AS num_posts,
        COUNT(comments.id) AS num_comments
    FROM likes
    LEFT JOIN posts ON posts.id = likes.post_id
    LEFT JOIN comments ON comments.id = likes.comment_id
    GROUP BY week
    ORDER BY week
) WITH DATA;
