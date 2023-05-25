CREATE TABLE infographic_banners (
    id int(10) unsigned NOT NULL AUTO_INCREMENT,
    title varchar(255) NOT NULL,
    order int(10) DEFAULT NULL,
    image json DEFAULT NULL,
    link varchar(255) DEFAULT NULL,
    is_active tinyint(1) DEFAULT 1,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE INDEX idx_infographic_banners_title ON infographic_banners (title);
CREATE INDEX idx_infographic_banners_is_active ON infographic_banners (is_active);
