CREATE TABLE IF NOT EXISTS "users" (
  "username" varchar(255) PRIMARY KEY,
  "hashed_password" varchar(255) NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT now(),
  "email" varchar(255) UNIQUE NOT NULL,
  "full_name" varchar(255) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now()
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

-- CREATE INDEX "accounts_index_1" ON "accounts" ("owner", "currency"); 
ALTER TABLE "accounts" ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner", "currency")
