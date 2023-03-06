ALTER TABLE masterdata_services
ADD COLUMN `status` VARCHAR(50)
DEFAULT "DRAFT"
AFTER `additional_information`;
