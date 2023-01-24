CREATE TABLE pop_up_banners (
    id int(10) NOT NULL AUTO_INCREMENT,
    title varchar(255) NOT NULL,
    button_label varchar(50) DEFAULT NULL,
    image json DEFAULT NULL,
    link varchar(255) DEFAULT NULL,
    status varchar(50) DEFAULT "NON-ACTIVE",
    duration int(10) DEFAULT NULL,
    start_date date,
    end_date date,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE INDEX idx_pop_up_banners_title ON pop_up_banners (title);
CREATE INDEX idx_pop_up_banners_status ON pop_up_banners (status);
