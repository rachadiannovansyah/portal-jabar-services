BEGIN;

DROP TABLE IF EXISTS `areas`;
CREATE TABLE `areas`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `depth` int(11) NULL DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `parent_code_kemendagri` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `code_kemendagri` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `code_bps` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `latitude` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `longitude` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `meta` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `areas_code_kemendagri_unique`(`code_kemendagri`) USING BTREE,
  UNIQUE INDEX `areas_code_bps_unique`(`code_bps`) USING BTREE,
  INDEX `areas_name_index`(`name`) USING BTREE,
  INDEX `areas_parent_code_kemendagri_index`(`parent_code_kemendagri`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

COMMIT;