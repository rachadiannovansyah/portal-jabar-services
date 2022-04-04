ALTER TABLE users ADD COLUMN status VARCHAR(20) NULL AFTER last_active;
CREATE INDEX idx_status ON users (status);