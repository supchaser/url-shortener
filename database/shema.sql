CREATE TABLE url(
		id         BIGINT      GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		alias 	   TEXT        NOT NULL UNIQUE,
		url   	   TEXT        NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_alias ON url(alias);
