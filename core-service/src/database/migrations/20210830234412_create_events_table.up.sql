BEGIN;

DROP TABLE IF EXISTS events;
CREATE TABLE `events` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `title` varchar(80) NOT NULL,
    `description` varchar(255),
    `date` date NOT NULL,
    `start_hour` varchar(18) NOT NULL,
    `end_hour` varchar(18) NOT NULL,
    `image` varchar(255) DEFAULT NULL,
    `published_by` varchar(16) DEFAULT NULL,
    `address` varchar(255) DEFAULT NULL,
    `category_id` int(10) unsigned NOT NULL,
    `province_code` varchar(191) unsigned NULL,
    `city_code` varchar(191) unsigned NULL,
    `district_code` varchar(191) unsigned NULL,
    `village_code` varchar(191) unsigned NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `events_categories_id_fk` (`category_id`),
    CONSTRAINT `events_categories_fk` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
    ) ENGINE=InnoDB DEFAULT CHARSET=latin1;
CREATE INDEX `idx_title` ON events (`title`);
CREATE INDEX `idx_start_hour` ON events (`start_hour`);
CREATE INDEX `idx_end_hour` ON events (`end_hour`);

COMMIT;