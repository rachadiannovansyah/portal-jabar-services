ALTER TABLE masterdata_services
DROP COLUMN updated_at,
DROP COLUMN created_at;

ALTER TABLE government_affairs
DROP COLUMN updated_at,
DROP COLUMN created_at;

ALTER TABLE spbe_rals
DROP COLUMN updated_at,
DROP COLUMN created_at;

ALTER TABLE uptd_cabdins
DROP COLUMN updated_at,
DROP COLUMN created_at;