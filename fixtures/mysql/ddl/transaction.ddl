SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS accounts;

CREATE TABLE
    IF not exists accounts (
        `sub` VARBINARY(255) NOT NULL,
        `user_id` VARBINARY(255) NOT NULL,
        `created_at` DATETIME NOT NULL PRIMARY KEY (`sub`)
    );