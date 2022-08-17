BEGIN;

CREATE TABLE IF NOT EXISTS "accounts" (
  "id" int PRIMARY KEY,
  "owner" varchar(255) NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar(255) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS "entries" (
  "id" int PRIMARY KEY,
  "account_id" bigint,
  "amount" bigint NOT NULL,
  "created_at" timestamptz
);

CREATE TABLE IF NOT EXISTS "transfers" (
  "id" int PRIMARY KEY,
  "from_account_id" int,
  "to_account_id" int,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL
);

CREATE INDEX "accounts_index_0" ON "accounts" ("owner");

CREATE INDEX "entries_index_1" ON "entries" ("account_id");

CREATE INDEX "transfers_index_2" ON "transfers" ("from_account_id");

CREATE INDEX "transfers_index_3" ON "transfers" ("to_account_id");

CREATE INDEX "transfers_index_4" ON "transfers" ("from_account_id", "to_account_id");

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" (id);

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" (id);

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" (id);

COMMIT;