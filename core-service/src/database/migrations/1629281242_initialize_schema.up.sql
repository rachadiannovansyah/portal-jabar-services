BEGIN;

DROP TABLE IF EXISTS categories;
CREATE TABLE `categories` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(80) NOT NULL,
  `description` varchar(255),
  `type` varchar(80),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

DROP TABLE IF EXISTS informations;
CREATE TABLE `informations` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `category_id` int(10) unsigned NOT NULL,
  `title` varchar(80) NOT NULL,
  `excerpt` varchar(150) NOT NULL,
  `content` text NOT NULL,
  `slug` varchar(100) DEFAULT NULL,
  `image` varchar(255) DEFAULT NULL,
  `source` varchar(80) DEFAULT NULL,
  `show_date` datetime NOT NULL,
  `end_date` datetime NOT NULL,
  `status` varchar(12) NOT NULL DEFAULT 'PUBLISHED',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `info_categories_id_fk` (`category_id`),
  CONSTRAINT `info_categories_id_fk` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
CREATE INDEX `idx_title` ON informations (`title`);
CREATE INDEX `idx_status` ON informations (`status`);
CREATE INDEX `idx_show_date` ON informations (`show_date`);
CREATE INDEX `idx_end_date` ON informations (`end_date`);

DROP TABLE IF EXISTS news;
CREATE TABLE `news` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `category_id` int(10) unsigned NOT NULL,
  `title` varchar(80) NOT NULL,
  `excerpt` varchar(150) NOT NULL,
  `content` text NOT NULL,
  `slug` varchar(100) DEFAULT NULL,
  `image` varchar(255) DEFAULT NULL,
  `video` varchar(80) DEFAULT NULL,
  `source` varchar(80) DEFAULT NULL,
  `status` varchar(12) NOT NULL DEFAULT 'PUBLISHED',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `news_categories_id_fk` (`category_id`),
  CONSTRAINT `news_categories_id_fk` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
CREATE INDEX `idx_title` ON news (`title`);
CREATE INDEX `idx_status` ON news (`status`);

-- ipj_db.units definition
CREATE TABLE `units` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(10) DEFAULT NULL,
  `name` varchar(100) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `logo` varchar(255) DEFAULT NULL,
  `website` varchar(60) DEFAULT NULL,
  `phone` varchar(100) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `chief` varchar(100) DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
CREATE INDEX `idx_name` ON units (`name`);

COMMIT;
