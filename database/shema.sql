CREATE TABLE url(
		id INTEGER PRIMARY KEY,
		alias TEXT NOT NULL UNIQUE,
		url TEXT NOT NULL);

CREATE INDEX idx_alias ON url(alias);
