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

--- adding it by create table method

CREATE TABLE users(
    email TEXT NOT NULL UNIQUE,
);