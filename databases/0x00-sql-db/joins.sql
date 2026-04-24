CREATE TABLE cars(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);


CREATE TABLE users(
    user_id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    car_id SERIAL REFERENCES cars(id)
);


--- innner joins -- joins two tables and create new column for refence column
--- returns rows that has a match


SELECT * FROM users
JOIN cars ON user_id = car_id;


---- left joins -- joins two tables list all records even if there is no match

SELECT * FROM users
LEFT JOIN cars ON users.user_id = cars.id;