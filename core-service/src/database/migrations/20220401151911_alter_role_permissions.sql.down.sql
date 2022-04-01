BEGIN;
ALTER TABLE roles DROP COLUMN label;
DROP INDEX idx_name ON roles;
DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS roles_permissions
COMMIT;