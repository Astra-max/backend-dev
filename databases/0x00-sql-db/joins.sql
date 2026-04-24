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


SELECT * FROM users
JOIN cars ON user_id = car_id;