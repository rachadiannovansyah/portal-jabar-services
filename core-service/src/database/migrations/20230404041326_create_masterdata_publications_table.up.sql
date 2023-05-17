CREATE TABLE `masterdata_publications` (
    `id` int(10) PRIMARY KEY AUTO_INCREMENT,
    `mds_id` int(10),
    `portal_category` varchar(150),
    `slug` varchar(255),
    `cover` json,
    `images` json,
    `infographics` json,
    `keywords` varchar(255),
    `faq` json,
    `status` varchar(100),
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (`mds_id`) REFERENCES `masterdata_services` (`id`)
);
