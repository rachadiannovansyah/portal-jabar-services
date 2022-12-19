ALTER TABLE events
DROP INDEX idx_category,
DROP INDEX idx_type,
DROP INDEX idx_date,
DROP INDEX idx_priority,
DROP INDEX idx_updated_at,
DROP INDEX idx_created_at,
DROP INDEX idx_deleted_at;