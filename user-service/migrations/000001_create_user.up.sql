CREATE TABLE IF NOT EXISTS users (
	id 		   VARCHAR(50) 				PRIMARY KEY,
	first_name VARCHAR(30) 				NOT NULL,
	email 	   VARCHAR(50) 				NOT NULL,
	password   VARCHAR(50) 				NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
	UNIQUE(email)
);
