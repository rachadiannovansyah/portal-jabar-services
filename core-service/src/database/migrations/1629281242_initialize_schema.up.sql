BEGIN;

CREATE TABLE `categories` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(80) NOT NULL,
  `description` text NOT NULL,
  `type` varchar(80) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `contents` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `categoryId` int(10) unsigned NOT NULL,
  `title` varchar(80) NOT NULL,
  `content` text NOT NULL,
  `slug` varchar(100) DEFAULT NULL,
  `image` varchar(255) DEFAULT NULL,
  `video` varchar(80) DEFAULT NULL,
  `source` varchar(80) DEFAULT NULL,
  `type` varchar(80) DEFAULT NULL,
  `showDate` datetime NOT NULL,
  `endDate` datetime NOT NULL,
  `status` varchar(12) NOT NULL DEFAULT 'PUBLISHED',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `categories_id_fk` (`categoryId`),
  CONSTRAINT `categories_id_fk` FOREIGN KEY (`categoryId`) REFERENCES `categories` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

COMMIT;