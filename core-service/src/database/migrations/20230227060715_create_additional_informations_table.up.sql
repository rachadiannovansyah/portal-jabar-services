CREATE TABLE `additional_informations` (
    `id` int(10) PRIMARY KEY AUTO_INCREMENT,
    `responsible_name` varchar(150),
    `phone_number` varchar(100),
    `email` varchar(150),
    `social_media` json
);
