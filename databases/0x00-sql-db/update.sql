--- updating all columns/dangerous in production

UPDATE users SET id = 90;

---- updating columns that has made certain condutions

UPDATE users SET id = 90 WHERE age > 34 AND age < 90;