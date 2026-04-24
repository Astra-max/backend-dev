-- primary key uniquely identifies a field

CREATE TABLE users(
    user_id SERIAL PRIMARY KEY
);

--- drop primary key constraint
--- attempting to return primary key after duplicated values throws an error


ALTER TABLE users DROP CONSTRAINT person_pkey;

--- adding primary key for fields
-- takes an array of columns to be made unique PRIMARY KEY(user_id, email)

ALTER TABLE users ADD PRIMARY KEY(user_id);

--- unique constrait allow us to have unique values for a given column
--- primary key ---> helps us identify unique fields
--- unique --> allows column to have unique values

--- adding it by create table method

CREATE TABLE users(
    email TEXT NOT NULL UNIQUE,
);

--- adding unique constraint via alter command

ALTER TABLE user ADD CONSTRAINT unique_email UNIQUE (email);

--- adding constraint via check key word

ALTER TABLE users ADD CONSTRAINT gender_pkey CHECK(gender='Female' or gender = 'Male');

--- connecting multiple table together via foreign key
--- foreign key is a ref that ref a primary key in another table
--- inorder to work the types have to be the same

CREATE TABLE person(
    id BIGSERIAL NOT NULL PRIMARY key,
    name TEXT NOT NULL,
    car_id BIGSERIAL REFERENCES UNIQUE car(id);
);

CREATE TABLE car(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name TEXT NOT NULL
)

--- UPDATING foregn key id column

UPDATE person SET car_id = 1 WHERE user_id = 1;


----- deleting columns with constraints
--- to delete a entry with constraint with delete the constraint first before the actual entry
--- the table the it is referenced
---- cascade remove every single row that is reference by that key

---- 1. first approach you delete manually which is safe

DELETE FROM user WHERE car_id = 1;
DELETE FROM cars WHERE id = 1;