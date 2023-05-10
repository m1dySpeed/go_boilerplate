CREATE TABLE "taxi_offices_archives"(
    "id" UUID NOT NULL,
    "city_id" UUID NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "latitude" DOUBLE PRECISION NOT NULL,
    "longitude" DOUBLE PRECISION NOT NULL,
    "working_hours" CHAR(255) NOT NULL,
    "officer_name" CHAR(255) NOT NULL,
    "phone" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "deleted_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "taxi_offices_archives" ADD PRIMARY KEY("id");
CREATE TABLE "cars"(
    "id" UUID NOT NULL,
    "office_id" UUID NOT NULL,
    "gov_number" VARCHAR(255) NOT NULL,
    "driver_surname" VARCHAR(255) NOT NULL,
    "driver_name" VARCHAR(255) NOT NULL,
    "driver_license" VARCHAR(255) NOT NULL,
    "add_driver_name" VARCHAR(255) NOT NULL,
    "add_driver_surname" VARCHAR(255) NOT NULL,
    "add_driver_license" VARCHAR(255) NOT NULL,
    "notice" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "cars" ADD PRIMARY KEY("id");
CREATE TABLE "orders"(
    "id" UUID NOT NULL,
    "car_id" UUID NOT NULL,
    "pickup_address_str" VARCHAR(255) NOT NULL,
    "destination_address_str" VARCHAR(255) NOT NULL,
    "pickup_latitude" DOUBLE PRECISION NOT NULL,
    "pickup_longitude" DOUBLE PRECISION NOT NULL,
    "destination_latitude" DOUBLE PRECISION NOT NULL,
    "destination_longitude" DOUBLE PRECISION NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "orders" ADD PRIMARY KEY("id");
CREATE TABLE "cities"(
    "id" UUID NOT NULL,
    "name" CHAR(255) NOT NULL,
    "latitude" DOUBLE PRECISION NOT NULL,
    "longitude" DOUBLE PRECISION NOT NULL
);
ALTER TABLE
    "cities" ADD PRIMARY KEY("id");
CREATE TABLE "taxi_offices"(
    "id" UUID NOT NULL,
    "city_id" UUID NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "latitude" DOUBLE PRECISION NOT NULL,
    "longitude" DOUBLE PRECISION NOT NULL,
    "working_hours" CHAR(255) NOT NULL,
    "officer_name" CHAR(255) NOT NULL,
    "phone" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "updated_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "taxi_offices" ADD PRIMARY KEY("id");
ALTER TABLE
    "taxi_offices" ADD CONSTRAINT "taxi_offices_city_id_foreign" FOREIGN KEY("city_id") REFERENCES "cities"("id");
ALTER TABLE
    "cars" ADD CONSTRAINT "cars_office_id_foreign" FOREIGN KEY("office_id") REFERENCES "taxi_offices"("id");
ALTER TABLE
    "orders" ADD CONSTRAINT "orders_car_id_foreign" FOREIGN KEY("car_id") REFERENCES "cars"("id");
ALTER TABLE
    "taxi_offices_archives" ADD CONSTRAINT "taxi_offices_archives_city_id_foreign" FOREIGN KEY("city_id") REFERENCES "cities"("id");