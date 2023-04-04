ALTER TABLE applications
ADD COLUMN `is_active` tinyint(1) AFTER features,
ADD COLUMN `title` varchar(255) AFTER is_active;
