ALTER TABLE news 
DROP INDEX news_start_date_index,
DROP INDEX news_end_date_index,
ADD INDEX news_start_date_index(views ASC) USING BTREE,
ADD INDEX news_end_date_index(views ASC) USING BTREE,
DROP INDEX news_updated_at_index,
DROP INDEX news_created_at_index,
DROP INDEX news_deleted_at_index;