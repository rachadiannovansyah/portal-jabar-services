BEGIN;
DROP TABLE IF EXISTS feedback;
CREATE TABLE feedback  (
  id bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  rating int(11) DEFAULT NULL,
  compliments text NOT NULL,
  criticism text NOT NULL,
  suggestions text NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);
COMMIT;
