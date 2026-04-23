-- primary key uniquely identifies a field

CREATE TABLE users(
    user_id SERIAL PRIMARY KEY
);

--- drop primary key constraint

ALTER TABLE users DROP CONSTRAINT person_pkey;