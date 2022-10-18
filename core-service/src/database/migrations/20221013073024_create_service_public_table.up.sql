CREATE TABLE service_public (
    id INT(10) NOT NULL AUTO_INCREMENT,
    general_information_id INT(10) UNSIGNED NOT NULL,
    purpose json,
    facility json,
    requirement json,
    tos json,
    info_graphic json,
    faq json,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (general_information_id) REFERENCES general_informations (id)
);
