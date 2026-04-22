-- count number of entries in a colunm

SELECT COUNT(age) FROM users;

-- finding a maximum value in a column

SELECT MAX(age) FROM users;

--- getting average of all values in  a column

SELECT AVG(age) FROM users;

--- selecting the minimum value within a column

SELECT MIN(age) FROM users;

--- add all values in a column

SELECT SUM(age) FROM users;

--- concatinating to a string column entry

SELECT user, STRING_AGG(user, "--> user") FROM users;

-- provide default value if not provided

SELECT user, COALESCE(email, "email not found") FROM users;
