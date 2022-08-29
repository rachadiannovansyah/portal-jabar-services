CREATE TABLE document_archives (
    id INT(10) NOT NULL AUTO_INCREMENT,
    title VARCHAR(150) NOT NULL,
    excerpt TEXT NOT NULL,
    description TEXT NOT NULL,
    source VARCHAR(80) DEFAULT NULL,
    mimetype VARCHAR(30) DEFAULT NULL,
    category VARCHAR(50) NOT NULL,
    year_published VARCHAR(30) DEFAULT NULL,
    created_by VARCHAR(36) DEFAULT NULL,
    updated_by VARCHAR(36) DEFAULT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
CREATE INDEX idx_title_documents ON document_archives (title);
CREATE INDEX idx_category_documents ON document_archives (category);
CREATE INDEX ids_year_published_documents ON document_archives (year_published);