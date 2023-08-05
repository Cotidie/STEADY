# PostgreSQL

## Challenges
- [ ] Writing efficienet queries
- [ ] Designing database (schema, structure)
- [ ] Understanding when to use advanced features
- [ ] Managing database

## SQL
To create a database, a schema should be defined for a table. It holds:
- What kind of thing to store
- What kind of properties it has
- What types of data they need

Let's create a table to store info about world's largest cities.  
The dataset can be found 🔗️[here](https://en.wikipedia.org/wiki/List_of_largest_cities).
```sql
CREATE TABLE cities(
    name VARCHAR(50),
    country VARCHAR(50),
    population INTEGER,
    area INTEGER
);
```
- **keyword**: (upper-case) reserved keywords for PostgreSQL
- **identifier**: (lower-case) user defined variables

### CREATE
```sql
INSERT INTO cities(name, country, population, area)
-- Multiple rows are separated by commas(,)
VALUES 
    ('Delhi', 'India', 28125000, 2240),
    ('Shanghai', 'China', 22125000, 4015),
    ('Sao Paulo', 'Brazil', 20935000, 3043),
    ('Tokyo', 'Japan', 38505000, 8223);
```

### READ
```sql
-- can use functions/operator to do calcauation
SELECT name, country, population / area AS density FROM cities;
```

### UPDATE

### DELETE