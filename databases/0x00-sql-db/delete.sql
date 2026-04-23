--- deleting all records from table

DELETE FROM users;

--- deleting a single record from a table

DELETE FROM users WHERE user_id = 23;

--- removing table entries that have satisfied the delete condition
-- via delete where and clause

DELETE FROM users WHERE gender = 'Male' AND country = 'Nigeria';