CREATE TABLE "user_info" (
  "user_id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "address" varchar NOT NULL,
  "balance" float NOT NULL
);

CREATE TABLE "book" (
  "book_id" bigserial PRIMARY KEY,
  "book_name" varchar NOT NULL,
  "stock" int NOT NULL,
  "price" float NOT NULL
);

CREATE TABLE "book_order" (
  "order_id" bigserial PRIMARY KEY,
  "user_name" varchar NOT NULL,
  "book_name" varchar NOT NULL,
  "amount" int NOT NULL,
  "address" varchar NOT NULL,
  "total_price" float NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "account" (
  "account_id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "email" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "user_info" ("user_id");

CREATE INDEX ON "book" ("book_id");

CREATE INDEX ON "book_order" ("order_id");

CREATE INDEX ON "account" ("account_id");

COMMENT ON COLUMN "book"."price" IS 'must be positive';

COMMENT ON COLUMN "book_order"."amount" IS 'must be positive';

COMMENT ON COLUMN "book_order"."total_price" IS 'must be positive';
