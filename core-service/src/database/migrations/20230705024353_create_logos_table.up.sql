CREATE TABLE logos (
    id int(10) unsigned NOT NULL AUTO_INCREMENT,
    title varchar(255) NOT NULL,
    image varchar(255) DEFAULT NULL,
    PRIMARY KEY (id)
);

CREATE INDEX idx_logos_title ON logos (title);
