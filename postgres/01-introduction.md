# PostgreSQL

## Installation
### PostgreSQL CLI
> Link: [Install psql](https://www.postgresql.org/download/)  

```bash
# 1. install `-contrib` package for additional utilities
sudo apt install postgresql postgresql-contrib
# 2. create a role associated with the current username or 'root'
sudo -u postgres createuser --interactive
----------------------------------------------
Output
Enter name of role to add: <username>
Shall the new role be a superuser? (y/n) y
----------------------------------------------
# 3. create a db with for the current username or 'root'
sudo -u postgres createdb <username>
# 4. `psql` command to run postgreSQL shell
sudo psql
# 5. list databases and switch to it
$ CREATE DATABASE <name>
$ \list
$ \connect <name>
```
 PostgreSQL authentication process requires both a role and a DB with the same name as UNIX account, or else login will fail. After successfully logging in to `psql shell`, you can type SQL commands for your liking.

### PostgreSQL GUI
> Link: [Install pgAdmin](https://www.pgadmin.org/download/pgadmin-4-apt/)  

```bash
# Install the public key for the repository (if not done previously):
curl -fsS https://www.pgadmin.org/static/packages_pgadmin_org.pub | sudo gpg --dearmor -o /usr/share/keyrings/packages-pgadmin-org.gpg

# Create the repository configuration file:
sudo sh -c 'echo "deb [signed-by=/usr/share/keyrings/packages-pgadmin-org.gpg] https://ftp.postgresql.org/pub/pgadmin/pgadmin4/apt/$(lsb_release -cs) pgadmin4 main" > /etc/apt/sources.list.d/pgadmin4.list && apt update'

# Install for both desktop and web modes:
sudo apt install pgadmin4

# Configure the webserver, if you installed pgadmin4-web:
sudo /usr/pgadmin4/bin/setup-web.sh
```

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
    id SERIAL PRIMARY KEY,
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
![phones-table](./images/01-phones-table.png)  

> - Write a query that will select the `name` of each phone and calculate the total revenue for each phone (`price` X `units_sold`)
> - Rename this calculated column to `revenue`
>- Filter out for all phones created by `Apple` or `Samsung`
```sql
SELECT name, price*units_sold as revenue
FROM phones
WHERE manufacturer IN ('Apple', 'Samsung');
```
- aliases (such as `revenue`) cannot be used in `WHERE`, `GROUP BY` clauses.

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

### INSERt
```sql
INSERT INTO <table>(<columns>)
VALUES (columns), (columns), ...;
```

## Operators
### String
- `||`, `CONCAT()`: Join two strings
- `LENGTH()`: Get the length of a string
- `UPPER()`, `LOWER()`: Make a string upper/lower cased
