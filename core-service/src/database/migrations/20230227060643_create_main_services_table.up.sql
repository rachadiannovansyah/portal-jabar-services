CREATE TABLE `main_services` (
    `id` int(10) PRIMARY KEY AUTO_INCREMENT,
    `information` int(10),
    `service_detail` int(10),
    `location` json,
    FOREIGN KEY (`information`) REFERENCES `service_informations` (`id`),
    FOREIGN KEY (`service_detail`) REFERENCES `service_details` (`id`)
);
