CREATE TABLE "User" (
  "Uuid" bigserial PRIMARY KEY,
  "FirstName" varchar(256) NOT NULL,
  "MiddleName" varchar(256) NOT NULL,
  "LastName" varchar(256) NOT NULL,
  "FullName" varchar NOT NULL GENERATED ALWAYS AS ("LastName" || ' ' || "FirstName" || ' ' || "MiddleName") STORED,
  "Gender" varchar(1) NOT NULL,
  "Age" smallint NOT NULL,
  "Balance" bigint NOT NULL DEFAULT 0,
  "Currency" varchar(3) NOT NULL
);

CREATE TABLE "Order" (
  "Uuid" bigserial PRIMARY KEY,
  "UserUuid" bigint NOT NULL,
  "Quantity" bigint NOT NULL
);

CREATE TABLE "Product" (
  "Uuid" bigserial PRIMARY KEY,
  "Description" varchar(256) NOT NULL,
  "Price" integer NOT NULL,
  "InStock" integer NOT NULL DEFAULT 0
);

CREATE TABLE "OrderProduct" (
  "OrderUuid" bigint NOT NULL,
  "ProductUuid" bigint NOT NULL
);

CREATE TABLE "UserToUser" (
  "FirstUserUuid" bigint NOT NULL ,
  "SecondUserUuid" bigint NOT NULL
);

CREATE INDEX ON "User" ("FullName");

CREATE INDEX ON "Product" ("Description");

CREATE INDEX ON "Order" ("UserUuid");

ALTER TABLE "OrderProduct" ADD FOREIGN KEY ("OrderUuid") REFERENCES "Order" ("Uuid") ON DELETE CASCADE;

ALTER TABLE "OrderProduct" ADD FOREIGN KEY ("ProductUuid") REFERENCES "Product" ("Uuid") ON DELETE CASCADE;

ALTER TABLE "UserToUser" ADD FOREIGN KEY ("FirstUserUuid") REFERENCES "User" ("Uuid") ON DELETE CASCADE;

ALTER TABLE "UserToUser" ADD FOREIGN KEY ("SecondUserUuid") REFERENCES "User" ("Uuid") ON DELETE CASCADE;