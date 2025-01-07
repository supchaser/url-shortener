CREATE TABLE "user"(
	u_id           BIGINT      GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	created_at     TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
	email          TEXT UNIQUE NOT NULL,
	password_hash  TEXT NOT NULL
);

CREATE TABLE url(
		id         BIGINT      GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		alias 	   TEXT        NOT NULL UNIQUE,
		url   	   TEXT        NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
		added_by   BIGINT,

		FOREIGN KEY (added_by) REFERENCES "user"(u_id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE INDEX idx_alias ON url(alias);

CREATE INDEX idx_user_email ON "user"(email);
