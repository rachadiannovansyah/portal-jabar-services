BEGIN;

DROP TABLE IF EXISTS tags_data;
CREATE TABLE tags_data (
    id int(10) unsigned NOT NULL AUTO_INCREMENT,
    data_id int(10) unsigned,
    tags_id int(10) unsigned,
    tags_name varchar(20),
    type varchar(10),
    PRIMARY KEY(id)
);
CREATE INDEX idx_tags_name ON tags_data (tags_name);

COMMIT;
