BEGIN;

DROP TABLE IF EXISTS agenda;
CREATE TABLE `agenda` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `title` varchar(80) NOT NULL,
    `description` varchar(255),
    `date` datetime NOT NULL,
    `start_hour` int(12) NOT NULL,
    `end_hour` int(12) NOT NULL,
    `image` varchar(255) DEFAULT NULL,
    `published_by` varchar(16) DEFAULT NULL,
    `address` varchar(80) DEFAULT NULL,
    `category_id` int(10) unsigned NOT NULL,
    `province_code` int(191) unsigned NOT NULL,
    `city_code` int(191) unsigned NOT NULL,
    `district_code` int(191) unsigned NOT NULL,
    `village_code` int(191) unsigned NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `agenda_categories_id_fk` (`category_id`),
    CONSTRAINT `agenda_categories_fk` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
    -- KEY `agenda_province_code_fk` (`province_code`),
    -- CONSTRAINT `agenda_province_code_fk` FOREIGN KEY (`province_code`) REFERENCES `areas` (`code_kemendagri`),
    -- KEY `agenda_city_code_fk` (`city_code`),
    -- CONSTRAINT `agenda_city_code_fk` FOREIGN KEY (`city_code`) REFERENCES `areas` (`code_kemendagri`),
    -- KEY `agenda_district_code_fk` (`district_code`),
    -- CONSTRAINT `agenda_district_code_fk` FOREIGN KEY (`district_code`) REFERENCES `areas` (`code_kemendagri`),
    -- KEY `agenda_village_code_fk` (`village_code`),
    -- CONSTRAINT `agenda_village_code_fk` FOREIGN KEY (`village_code`) REFERENCES `areas` (`code_kemendagri`)
    ) ENGINE=InnoDB DEFAULT CHARSET=latin1;
CREATE INDEX `idx_title` ON agenda (`title`);
CREATE INDEX `idx_start_hour` ON agenda (`start_hour`);
CREATE INDEX `idx_end_hour` ON agenda (`end_hour`);

COMMIT;