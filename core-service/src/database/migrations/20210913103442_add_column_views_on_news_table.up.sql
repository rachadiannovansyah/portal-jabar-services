BEGIN;
ALTER TABLE news
ADD views bigint DEFAULT 0 NOT NULL AFTER `status`;
CREATE INDEX news_views_index ON news (views);
COMMIT;
