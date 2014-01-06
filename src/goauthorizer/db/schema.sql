CREATE TABLE "sites" (
    "id" SERIAL UNIQUE,
    "domain" TEXT NOT NULL,
    "title" TEXT
);
CREATE UNIQUE CONSTRAINT "sites_domain_lower_uniq" ON "sites" (LOWER("domain"));


CREATE TABLE "accounts" (
    "id" SERIAL UNIQUE,
    "created" TIMESTAMP NOT NULL,
    "deleted" BOOLEAN NOT NULL DEFAULT FALSE
);


CREATE TABLE "token_email" (
    "id" INTEGER NOT NULL REFERENCES "accounts" ("id")
        ON UPDATE CASCADE ON DELETE CASCADE,
    "email" TEXT NOT NULL,
    "password" BYTEA
);
CREATE UNIQUE CONSTRAINT "email_lower_uniq" ON "token_email" (LOWER("email"));


CREATE TABLE "token_facebook" (
    "id" INTEGER NOT NULL REFERENCES "accounts" ("id")
        ON UPDATE CASCADE ON DELETE CASCADE,
    "external_id" TEXT NOT NULL UNIQUE
);
