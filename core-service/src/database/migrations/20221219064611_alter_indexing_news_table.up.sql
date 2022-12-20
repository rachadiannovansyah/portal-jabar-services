ALTER TABLE news
DROP INDEX news_start_date_index,
DROP INDEX news_end_date_index,
ADD INDEX news_start_date_index(start_date ASC) USING BTREE,
ADD INDEX news_end_date_index(end_date ASC) USING BTREE,
ADD INDEX news_updated_at_index(updated_at ASC) USING BTREE,
ADD INDEX news_created_at_index(created_at ASC) USING BTREE,
ADD INDEX news_deleted_at_index(deleted_at ASC) USING BTREE;
