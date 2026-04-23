---- handling constraint error
--- through duplicate value violents unique constraints
INSERT INTO users(id, name) VALUES(1, 'astra');

---- handling error

INSERT INTO users(id, name) VALUES(1, 'astra') ON CONFLICT (id) DO NOTHING;