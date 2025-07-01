package sqldb

const schema = `
CREATE TABLE IF NOT EXISTS users (
	uuid UUID PRIMARY KEY,
	username TEXT NOT NULL,
	email TEXT NOT NULL,
	password BYTEA NOT NULL,
	role TEXT NOT NULL,
	deleted TIMESTAMP,
	created TIMESTAMP NOT NULL,
	last_login TIMESTAMP
);

CREATE TABLE IF NOT EXISTS templates (
	uuid UUID PRIMARY KEY,
	owner_uuid UUID NOT NULL REFERENCES users(uuid),
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	preview TEXT,
	template TEXT NOT NULL,
	created TIMESTAMP NOT NULL,
	deleted TIMESTAMP
);

CREATE TABLE IF NOT EXISTS template_parameters (
	uuid UUID PRIMARY KEY,
	template_uuid UUID NOT NULL REFERENCES templates(uuid),
	name TEXT NOT NULL,
	type TEXT NOT NULL,
	display_name TEXT NOT NULL,
	default_value TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS parameter_constraint_types (
	constraint_type_id INT PRIMARY KEY,
	constraint_type_name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS template_parameters_constraints (
	uuid UUID PRIMARY KEY,
	template_parameter_uuid UUID NOT NULL REFERENCES template_parameters(uuid),
	constraint_type_id INT NOT NULL REFERENCES parameter_constraint_types(constraint_type_id),
	constraint_value TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS cache (
	key TEXT PRIMARY KEY,
	value BYTEA NOT NULL,
	created TIMESTAMP NOT NULL
);
`
