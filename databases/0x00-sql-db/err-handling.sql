---- handling constraint error
--- through duplicate value violents unique constraints
INSERT INTO users(id, name) VALUES(1, 'astra');

---- handling error
--- columns must be unique to use it

INSERT INTO users(id, name) VALUES(1, 'astra') ON CONFLICT (id) DO NOTHING;

---- update the current results to incomming req instead of nothing
--- EXCLUDED replaces the current field throwing constraint violation err

INSERT INTO users(id, name) VALUES(1, 'astra') ON CONFLICT (id) DO UPDATE SET id = EXCLUDED.id;
