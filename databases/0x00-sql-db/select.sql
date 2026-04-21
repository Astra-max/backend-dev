-- selecting all columns from a table

SELECT * FROM user;

--- select specific columns, id,name,age and email

SELECT user_id, first_name, age, email FROM users;

-- select columns with distinct value/not other occurences

SELECT DISTINCT first_name FROM users;

-- selecting table columns ordering by age in descending order

SELECT * FROM user ORDER BY age DESC;

-- selecting a limited number of columns

SELECT * FROM user LIMIT 10;