ALTER TABLE events
ADD type ENUM('offline', 'online') NOT NULL AFTER published_by;