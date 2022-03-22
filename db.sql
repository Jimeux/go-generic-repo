DROP DATABASE IF EXISTS generic_db;
CREATE DATABASE generic_db;

USE generic_db;

CREATE TABLE job
(
    id     INTEGER PRIMARY KEY AUTO_INCREMENT,
    name   varchar(200) not null
);

CREATE TABLE person
(
    id          INTEGER PRIMARY KEY AUTO_INCREMENT,
    given_name  varchar(200) not null,
    family_name varchar(200) not null
);
