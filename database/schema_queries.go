package database

var createEntitiesTableQuery = `
CREATE TABLE IF NOT EXISTS entities (
	id int UNIQUE NOT NULL AUTO_INCREMENT,
	name varchar(255) NOT NULL,
	description varchar(2055) NOT NULL,
	kind varchar(255) NOT NULL,
	importance int NOT NULL,
	status int NOT NULL DEFAULT '1',
	dateCreated TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
)
`
