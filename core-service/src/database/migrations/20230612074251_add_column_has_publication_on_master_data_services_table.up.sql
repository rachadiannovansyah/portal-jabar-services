ALTER TABLE masterdata_services
ADD COLUMN has_publication tinyint(1) DEFAULT 0 AFTER status;
