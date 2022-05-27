CREATE TABLE public_services (
    id INT(10) NOT NULL AUTO_INCREMENT,
    name varchar(100) NOT NULL,
    description text NULL,
    unit varchar(100) NULL,
    url varchar(255) NULL,
    image varchar(100) DEFAULT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
CREATE INDEX idx_public_service_name ON public_services (name);