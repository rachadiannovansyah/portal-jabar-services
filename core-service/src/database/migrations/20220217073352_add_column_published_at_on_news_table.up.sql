BEGIN;
ALTER TABLE news
ADD published_at timestamp AFTER `updated_by`;
CREATE INDEX news_published_at_index ON news (published_at);
COMMIT;