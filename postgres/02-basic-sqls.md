## SQL
To create a database, a schema should be defined for a table. It holds:
- What kind of thing to store
- What kind of properties it has
- What types of data they need

### Challenges
- [ ] Writing efficienet queries
- [ ] Designing database (schema, structure)
- [ ] Understanding when to use advanced features
- [ ] Managing database


### CREATE
Let's create a table to store info about world's largest cities.  
The dataset can be found 🔗️[here](https://en.wikipedia.org/wiki/List_of_largest_cities).
```sql
CREATE TABLE cities(
    id SERIAL PRIMARY KEY,
    name VARCHAR(50),
    country VARCHAR(50),
    population INTEGER,
    area INTEGER
);
```
- **keyword**: (upper-case) reserved keywords for PostgreSQL
- **identifier**: (lower-case) user defined variables

### READ
![phones-table](./images/01-phones-table.png)  

> - Write a query that will select the `name` of each phone and calculate the total revenue for each phone (`price` X `units_sold`)
> - Rename this calculated column to `revenue`
>- Filter out for all phones created by `Apple` or `Samsung`
```sql
SELECT (DISTINCT) name, price*units_sold as revenue
FROM phones
WHERE manufacturer IN ('Apple', 'Samsung')
ORDER BY name
LIMIT 2      -- display up to 2 rows
OFFSET 5;      -- skips first 5 rows
```
- aliases (such as `revenue`) cannot be used in `WHERE`, `GROUP BY` clauses.
- `DISTINCT` keyword only shows the rows with unique values

### UPDATE
```sql
UPDATE <table>
SET <expression>
WHERE ...
```

### DELETE
```sql
DELETE FROM <table>
WHERE ...
```

### INSERT
```sql
INSERT INTO <table>(<columns>)
VALUES (columns), (columns), ...;
```

## Operators
### String
- `||`, `CONCAT()`: Join two strings
- `LENGTH()`: Get the length of a string
- `UPPER()`, `LOWER()`: Make a string upper/lower cased
- `GREATEST()`, `LEAST()` can clip a value of each row with upper/lower bound.
