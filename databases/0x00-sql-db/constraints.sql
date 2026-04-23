-- primary key uniquely identifies a field

CREATE TABLE users(
    user_id SERIAL PRIMARY KEY
);

--- drop primary key constraint
--- attempting to return primary key after duplicated values throws an error


ALTER TABLE users DROP CONSTRAINT person_pkey;

--- adding primary key for fields

ALTER TABLE users ADD PRIMARY KEY(user_id);