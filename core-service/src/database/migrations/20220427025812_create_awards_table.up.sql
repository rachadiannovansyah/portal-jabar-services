CREATE TABLE awards (
    id INT(10) NOT NULL AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    logo VARCHAR(50) DEFAULT NULL,
    appreciator VARCHAR(100) DEFAULT NULL,
    description TEXT NOT NULL,
    category VARCHAR(50) NOT NULL,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
CREATE INDEX idx_title ON awards (title);
CREATE INDEX idx_category ON awards (category);