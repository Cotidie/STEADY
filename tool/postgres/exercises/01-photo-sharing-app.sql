-- 1. JOIN
-- For each comments, show the contents of the comment and the username of the user who wrote the comment
SELECT c.contents AS contents, u.username AS username
FROM comments as c
JOIN users AS u ON c.user_id = u.id;