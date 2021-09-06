BEGIN;

ALTER TABLE events
ADD `priority` ENUM('low', 'medium', 'high') DEFAULT 'low' NOT NULL AFTER `date`;

COMMIT;