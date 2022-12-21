-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2022-12-06T08:48:33.842Z

CREATE TABLE "Users" (
  "id" serial PRIMARY KEY,
  "email" varchar(50),
  "username" varchar(12),
  "password" varchar(20)
);

CREATE TABLE "UserDetails" (
  "id" serial PRIMARY KEY,
  "user_id" integer,
  "fullname" varchar(50),
  "phone" varchar(13),
  "gender" varchar(20)
);

ALTER TABLE "UserDetails" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("id");
