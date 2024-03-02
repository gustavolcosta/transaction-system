CREATE TABLE IF NOT EXISTS accounts (id SERIAL PRIMARY KEY,document_number VARCHAR(255));

CREATE TABLE IF NOT EXISTS operations_types (id SERIAL PRIMARY KEY, description VARCHAR(255), negative_amount BOOLEAN);