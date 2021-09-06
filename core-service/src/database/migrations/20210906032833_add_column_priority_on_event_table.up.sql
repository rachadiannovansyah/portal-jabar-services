BEGIN;

ALTER TABLE events
ADD `priority` int(80) NOT NULL AFTER `date`;

COMMIT;