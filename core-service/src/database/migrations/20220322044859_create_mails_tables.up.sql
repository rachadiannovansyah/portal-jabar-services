CREATE TABLE mails (
    id int(10) NOT NULL AUTO_INCREMENT,
    sender varchar(80) NOT NULL,
    receiver varchar(80) NOT NULL,
    subject varchar(255) NOT NULL,
    cc varchar(80),
    body varchar(255),
    template varchar(50) NOT NULL,
    PRIMARY KEY (id)
);
CREATE INDEX idx_template ON mails (template);