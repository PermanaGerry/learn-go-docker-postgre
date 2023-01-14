-- install extension
CREATE EXTENSION "uuid-ossp";

-- create table
CREATE TABLE IF NOT EXISTS items (
    id uuid DEFAULT uuid_generate_v4(),
    name VARCHAR NOT NULL,
	stock int NOT NULL,

    PRIMARY KEY (id)
);