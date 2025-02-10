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
	preview BYTEA NOT NULL,
	template TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS template_parameters (
	uuid UUID PRIMARY KEY,
	template_uuid UUID NOT NULL REFERENCES templates(uuid),
	name TEXT NOT NULL,
	type TEXT NOT NULL,
	display_name TEXT NOT NULL,
	default_value TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS template_parameters_constrains_types (
	id INT PRIMARY KEY,
	name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS template_parameters_constrains (
	uuid UUID PRIMARY KEY,
	template_parameter_uuid UUID NOT NULL REFERENCES template_parameters(uuid),
	constrain_type INT NOT NULL REFERENCES template_parameters_constrains_types(id),
	constrain_value TEXT NOT NULL
);
`
