# Database and Table
![database-design](./images/02-database-design.png)  
Design a database of minimized Instagram having features such as:
  - Users have followers and can follow other users
  - Users can like a photo
  - Users have a profile page and post a list of photos
  - comments are created by users

## Components
### Relationship
![relationships](./images/03-relationship-arrows.png)  
Relationships are a basic building block between tables. It's important to state a relationship in the perspective of one table, such as `A user has multiple photos, and a photo has only one user`.

Associated two tables can be read using `JOIN` keyword
```sql
-- `username` belongs to `users` table
SELECT url, username FROM photos
JOIN users ON users.id = photos.user_id;
```

## Keys
```sql
CREATE TABLE cities(
    id SERIAL PRIMARY KEY,
    name VARCHAR(50),
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE
);
```
- **Primary Key** : Uniquely identifies a record in the table
    - may use an auto generated IDs using `SERIAL` keyword
- **Foreign Key** : Identifies a record that a row is associated with
    - use keyword `REFERENCES <table>(<column>)` to point at some foreign table
  
## Constraints
`Constraint` is defined on foreign keys to deal with changes with the foreign table's records being referenced. Below are its options:
- `ON DELETE NO ACTION`: by default. throws an error
- `ON DELETE RESTRICT`:  throws an error
- `ON DELETE CASCADE`: delete the associated records, too
- `ON DELETE SET NULL`: set the foreign key column's value to 'NULL'
- `ON DELETE SET DEFAULT`: set the foreign key column's value to the default 

## Join
 `JOIN` produces values by merging rows together from different tables having relationships, mostly by foreign key references. Below are SQL use-cases using JOIN.  

 ### Types
![joins](./images/04-four-kinds-of-joins.png)  
#### Inner Join
```sql
SELECT * FROM phones
JOIN users ON users.is = phones.user_id;
```
Results in only the records matching the keys that exist both the table `users` and `phones`. Other records are discarded.

#### LEFT/RIGHT OUTER JOIN
```sql
SELECT * FROM phones
LEFT/RIGHT JOIN users ON users.id = phones.user_id;
```
Results in printing all the records on the left / right table and set to NULL on the other columns not matching the foreign keys. In this case, `phones` (the table selecting from) is to be `LEFT` and `users` to be `RIGHT`.


 ### Exercises
 (run `photo-sharing-app-data.sql` to set up and populate tables)
```sql
-- For each comments, show the contents of the comment and the username of the user who wrote the comment
SELECT c.contents AS contents, u.username AS username
FROM comments as c
JOIN users AS u ON c.user_id = u.id;
```

## Group
![grouped](./images/05-grouped-picturing.png)  
 Grouping reduces rows down to fewer rows using the `GROUP BY` keyword. Ungrouped columns cannot be `SELECT`ed or else such columns should be applied with `aggregates` functions.
 - **aggregates**: `COUNT`, `SUM`, `AVG`, `MIN`, `MAX`, ...
 - `COUNT(*)`: counts the number of all rows 
 - priorities between operations are: `FROM` > `JOIN` > `WHERE` > `GROUP BY` > `HAVING`
```sql
-- Given a table of phones, print the names of manufacturers and total revenue (price * units_sold) for all phones.  Only print the manufacturers who have revenue greater than 2,000,000 for all the phones they sold.
SELECT manufacturer, SUM(price * units_sold)
FROM phones
GROUP BY manufacturer
HAVING SUM(price * units_sold) > 2000000;
```


