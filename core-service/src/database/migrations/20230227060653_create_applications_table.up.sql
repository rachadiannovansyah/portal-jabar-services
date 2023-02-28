CREATE TABLE `applications` (
    `id` int(10) PRIMARY KEY AUTO_INCREMENT,
    `name` varchar(255),
    `status` varchar(50),
    `features` json
);
