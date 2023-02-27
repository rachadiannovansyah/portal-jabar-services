CREATE TABLE `masterdata_services` (
    `id` int(10) PRIMARY KEY AUTO_INCREMENT,
    `main_service` int(10),
    `application` int(10),
    `additional_information` int(10),
    FOREIGN KEY (`id`) REFERENCES `main_services` (`id`),
    FOREIGN KEY (`application`) REFERENCES `applications` (`id`),
    FOREIGN KEY (`additional_information`) REFERENCES `additional_informations` (`id`)
);
