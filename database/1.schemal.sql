CREATE SCHEMA internal;

USE internal;

CREATE DATABASE IF NOT EXISTS authentication
    CHARACTER SET = utf8
    COLLATE = utf8mb4
    ENCRYPTION = 'Y';

USE `authentication`;

CREATE TABLE IF NOT EXISTS `user`
(
    user_id BINARY(16) default (UUID_TO_BIN(UUID())) PRIMARY KEY,
    `name` VARCHAR(100) NOT NULL,
    country VARCHAR(3) NOT NULL,
    `enabled` bit DEFAULT 1
);

CREATE TABLE IF NOT EXISTS user_login
(
    uuid BINARY(16) default (UUID_TO_BIN(UUID())),
    user_id BINARY(16) default (UUID_TO_BIN(UUID())),
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    PRIMARY KEY (uuid),
    FOREIGN KEY (user_id) REFERENCES `user`(user_id),
    UNIQUE INDEX (email)
);