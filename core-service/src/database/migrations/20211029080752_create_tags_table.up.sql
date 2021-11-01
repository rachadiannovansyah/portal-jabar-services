BEGIN;

DROP TABLE IF EXISTS tags;
CREATE TABLE tags (
    id int(10) unsigned NOT NULL AUTO_INCREMENT,
    name varchar(20) NOT NULL,
    PRIMARY KEY(id)
);
CREATE INDEX idx_tags_name ON tags (name);

COMMIT;