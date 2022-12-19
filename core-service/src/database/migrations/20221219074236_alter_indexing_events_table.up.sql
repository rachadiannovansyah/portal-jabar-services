ALTER TABLE events
ADD INDEX idx_category(category ASC) USING BTREE,
ADD INDEX idx_type(type ASC) USING BTREE,
ADD INDEX idx_date(date ASC) USING BTREE,
ADD INDEX idx_priority(priority ASC) USING BTREE,
ADD INDEX idx_updated_at(updated_at ASC) USING BTREE,
ADD INDEX idx_created_at(created_at ASC) USING BTREE,
ADD INDEX idx_deleted_at(deleted_at ASC) USING BTREE;