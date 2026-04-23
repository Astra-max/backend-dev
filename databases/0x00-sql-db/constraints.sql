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