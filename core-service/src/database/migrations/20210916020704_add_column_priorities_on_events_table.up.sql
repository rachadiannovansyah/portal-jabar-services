BEGIN;

ALTER TABLE events
ADD `priorities` int(10) DEFAULT 1 NOT NULL AFTER `date`;

COMMIT;