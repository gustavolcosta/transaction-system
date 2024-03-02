CREATE TABLE IF NOT EXISTS accounts (id SERIAL PRIMARY KEY,document_number VARCHAR(255));

CREATE TABLE IF NOT EXISTS operations_types (id SERIAL PRIMARY KEY, description VARCHAR(255), negative_amount BOOLEAN);

CREATE TABLE IF NOT EXISTS transactions (id SERIAL PRIMARY KEY, account_id INT REFERENCES accounts(id), operation_type_id INT REFERENCES operations_types(id), amount FLOAT, event_date TIMESTAMP);