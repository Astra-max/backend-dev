-- selecting all columns from a table

SELECT * FROM user;

--- select specific columns, id,name,age and email

SELECT user_id, first_name, age, email FROM users;

-- select columns with distinct value/not other occurences

SELECT DISTINCT first_name FROM users;