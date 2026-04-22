-- selecting all columns from a table

SELECT * FROM user;

--- select specific columns, id,name,age and email

SELECT user_id, first_name, age, email FROM users;

-- select columns with distinct value/not other occurences

SELECT DISTINCT first_name FROM users;

-- selecting table columns ordering by age in descending order

SELECT * FROM user ORDER BY age DESC;

-- organizing data in ascending order

SELECT age FROM users ORDER BY age ASC;

-- selecting a limited number of columns

SELECT * FROM user LIMIT 10;

-- selecting column based on age number

SELECT * FROM users WHERE age == 4;

-- selecting row in a particular range

SELECT age FROM users WHERE age > 10 AND age < 20;

-- or clause

SELECT * FROM users WHERE age = 20 OR age < 40;

-- select row wher a pattern is matched

SELECT email FROM users WHERE email ILIKE "%o@gmail.com"