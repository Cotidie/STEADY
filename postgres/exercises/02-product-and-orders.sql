-- 1. Subquery
-- Calculate the average price of phones for each manufacturer.  Then print the highest average price. Rename this value to max_average_price
SELECT MAX(average_price) AS max_average_price
FROM (
    SELECT AVG(price) AS average_price
    FROM phones
    GROUP BY manufacturer
) AS p

-- 2. ALL/SOME with subquery
-- Write a query that prints the name of all phones that have a price greater than all phones made by Samsung.
SELECT name
FROM phones
WHERE price > ALL (
    SELECT price
    FROM phones
    WHERE manufacturer = 'Samsung'
);

-- 3. Correlated Subquery
-- Print the number of orders for each product without using join or a group by
SELECT name, (
    SELECT COUNT(*)
    FROM products AS o1
    WHERE o1.product_id = p1.id
)
FROM products AS p1;

-- 4. DISTINCT
-- Write a query that will print the number of unique phone manufacturers.
SELECT COUNT(DISTINCT manufacturer)
FROM phones;