CREATE TABLE "Garage" (
  "id" bigserial PRIMARY KEY,
  "car_id" bigint,
  "brand_id" bigint,
  "owner_id" bigint
);

CREATE TABLE "Cars" (
  "id" bigserial PRIMARY KEY,
  "Make" varchar,
  "Model" varchar,
  "Horsepower" bigint,
  "Bought_at" timestamptz DEFAULT 'now()'
);

CREATE TABLE "Brand" (
  "id" bugserial PRIMARY KEY,
  "Name" varchar,
  "Country" varchar
);

CREATE TABLE "Owner" (
  "id" bigserial PRIMARY KEY,
  "FirstName" varchar,
  "LastName" varchar,
  "Age" int,
  "Address" varchar
);

ALTER TABLE "Garage" ADD FOREIGN KEY ("car_id") REFERENCES "Cars" ("id");

ALTER TABLE "Garage" ADD FOREIGN KEY ("brand_id") REFERENCES "Cars" ("Make");

ALTER TABLE "Garage" ADD FOREIGN KEY ("owner_id") REFERENCES "Owner" ("id");

ALTER TABLE "Cars" ADD FOREIGN KEY ("Make") REFERENCES "Brand" ("Name");
