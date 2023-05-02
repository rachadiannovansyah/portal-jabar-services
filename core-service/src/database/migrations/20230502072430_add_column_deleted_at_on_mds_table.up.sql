ALTER TABLE masterdata_services
ADD COLUMN deleted_at timestamp NULL DEFAULT NULL AFTER created_at;