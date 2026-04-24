--- creating a database to user

CREATE DATABASE westmart OWNER astra;
CREATE USER astra WITH PASSWORD '123';

GRANT ALL PRIVILEGES ON DATABASE westmart TO astra;