CREATE TABLE `service_details` (
  `id` int(10) PRIMARY KEY AUTO_INCREMENT,
  `terms_and_condition` varchar(255),
  `service_procedures` varchar(255),
  `service_fee` varchar(30),
  `operational_time` json,
  `hotline_number` varchar(30),
  `hotline_mail` varchar(100)
);
