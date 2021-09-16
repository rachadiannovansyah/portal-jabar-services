ALTER TABLE events
ADD priorities tinyint(1) DEFAULT 1 NOT NULL AFTER date;