CREATE TABLE districts (
    id INT(10) NOT NULL AUTO_INCREMENT,
    name varchar(100) NOT NULL,
    chief varchar(100)NULL,
    address varchar(150) NULL,
    website varchar(100) NULL,
    logo varchar(255) DEFAULT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
CREATE INDEX idx_districts_name ON districts (name);