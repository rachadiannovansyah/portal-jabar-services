CREATE TABLE quick_accesses (
    id int(10) unsigned NOT NULL AUTO_INCREMENT,
    title varchar(255) NOT NULL,
    description varchar(255) NOT NULL,
    image varchar(255) DEFAULT NULL,
    link varchar(255) DEFAULT NULL,
    is_active tinyint(1) DEFAULT 0,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE INDEX idx_quick_accesses_title ON quick_accesses (title);
