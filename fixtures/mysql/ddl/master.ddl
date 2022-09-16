DROP TABLE IF EXISTS prefectures;

CREATE TABLE
    `prefectures` (
        `id` TINYINT UNSIGNED NOT NULL AUTO_INCREMENT,
        `name` VARCHAR(255) NOT NULL,
        PRIMARY KEY (`id`)
    );

DROP TABLE IF EXISTS products;

CREATE TABLE
    `products` (
        `id` TINYINT UNSIGNED NOT NULL AUTO_INCREMENT,
        `name` VARCHAR(255) NOT NULL,
        PRIMARY KEY (`id`)
    );