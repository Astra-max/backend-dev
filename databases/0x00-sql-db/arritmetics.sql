-- multiplying column value

SELECT age, age * 20 AS new_age FROM users;

-- basic multiplication arrithmetics

SELECT 7 * 7;

--- handling zero division error

SELECT COALESCE(10 / NULLIF(0, 0), "zero division erro");