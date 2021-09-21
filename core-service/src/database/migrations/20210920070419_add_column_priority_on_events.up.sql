ALTER TABLE events
ADD priority tinyint(1) DEFAULT 1 NOT NULL AFTER date;