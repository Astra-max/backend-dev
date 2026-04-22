-- substracting dte value

SELECT NOW() - INTERVAL '1 YEAR';

-- EXTRACTING values from a date

SELECT EXCTRACT(DAY, NOW());

-- getting age from current to the target

SELECT age AS current_age EXCTRACT(YEAR, AGE(NOW(), age)) AS years_taken FROM users;