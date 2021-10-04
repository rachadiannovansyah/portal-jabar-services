BEGIN;

DROP TABLE IF EXISTS categories;
CREATE TABLE categories (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  title varchar(80) NOT NULL,
  description varchar(255),
  type varchar(80),
  PRIMARY KEY (id)
);

DROP TABLE IF EXISTS informations;
CREATE TABLE informations (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  category_id int(10) unsigned NOT NULL,
  title varchar(80) NOT NULL,
  excerpt varchar(150) NOT NULL,
  content text NOT NULL,
  slug varchar(100) DEFAULT NULL,
  image varchar(255) DEFAULT NULL,
  source varchar(80) DEFAULT NULL,
  show_date datetime NOT NULL,
  end_date datetime NOT NULL,
  status varchar(12) NOT NULL DEFAULT 'PUBLISHED',
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  KEY info_categories_id_fk (category_id),
  CONSTRAINT info_categories_id_fk FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX idx_title ON informations (title);
CREATE INDEX idx_status ON informations (status);
CREATE INDEX idx_show_date ON informations (show_date);
CREATE INDEX idx_end_date ON informations (end_date);

DROP TABLE IF EXISTS news;
CREATE TABLE news (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  category_id int(10) unsigned NOT NULL,
  title varchar(80) NOT NULL,
  excerpt varchar(150) NOT NULL,
  content text NOT NULL,
  slug varchar(100) DEFAULT NULL,
  image varchar(255) DEFAULT NULL,
  video varchar(80) DEFAULT NULL,
  source varchar(80) DEFAULT NULL,
  status varchar(12) NOT NULL DEFAULT 'PUBLISHED',
  views bigint DEFAULT 0 NOT NULL,
  highlight tinyint(1) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  KEY news_categories_id_fk (category_id),
  CONSTRAINT news_categories_id_fk FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX idx_title ON news (title);
CREATE INDEX idx_status ON news (status);
CREATE INDEX news_views_index ON news (views);

-- ipj_db.units definition
DROP TABLE IF EXISTS units;
CREATE TABLE units (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  parent_id int(10) DEFAULT NULL,
  name varchar(100) NOT NULL,
  description varchar(255) DEFAULT NULL,
  logo varchar(255) DEFAULT NULL,
  website varchar(60) DEFAULT NULL,
  phone varchar(100) DEFAULT NULL,
  address varchar(255) DEFAULT NULL,
  chief varchar(100) DEFAULT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
CREATE INDEX idx_name ON units (name);

DROP TABLE IF EXISTS areas;
CREATE TABLE areas (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  depth int(11) NULL DEFAULT NULL,
  name varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  parent_code_kemendagri varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  code_kemendagri varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  code_bps varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  latitude varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  longitude varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  meta varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  created_at timestamp(0) NULL DEFAULT NULL,
  updated_at timestamp(0) NULL DEFAULT NULL,
  PRIMARY KEY (id) USING BTREE,
  UNIQUE INDEX areas_code_kemendagri_unique(code_kemendagri) USING BTREE,
  UNIQUE INDEX areas_code_bps_unique(code_bps) USING BTREE,
  INDEX areas_name_index(name) USING BTREE,
  INDEX areas_parent_code_kemendagri_index(parent_code_kemendagri) USING BTREE
);

DROP TABLE IF EXISTS events;
CREATE TABLE events (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  title varchar(80) NOT NULL,
  description varchar(255),
  date date NOT NULL,
  priority tinyint(1) DEFAULT 1 NOT NULL,
  start_hour varchar(18) NOT NULL,
  end_hour varchar(18) NOT NULL,
  image varchar(255) DEFAULT NULL,
  published_by varchar(16) DEFAULT NULL,
  type ENUM('offline', 'online') NOT NULL,
  address varchar(255) DEFAULT NULL,
  url varchar(80) DEFAULT NULL,
  category_id int(10) unsigned NOT NULL,
  province_code varchar(191) NULL,
  city_code varchar(191) NULL,
  district_code varchar(191) NULL,
  village_code varchar(191) NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  KEY events_categories_id_fk (category_id),
  CONSTRAINT events_categories_fk FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX idx_title ON events (title);
CREATE INDEX idx_start_hour ON events (start_hour);
CREATE INDEX idx_end_hour ON events (end_hour);

DROP TABLE IF EXISTS feedback;
CREATE TABLE feedback (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  rating int(11) DEFAULT NULL,
  compliments text NOT NULL,
  criticism text NOT NULL,
  suggestions text NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

DROP TABLE IF EXISTS featured_programs;
CREATE TABLE featured_programs (
	id int(10) unsigned NOT NULL AUTO_INCREMENT,
	title varchar(100) not null,
	excerpt varchar(255) not null,
	description text not null,
	organization varchar(100),
	categories json,
	service_type varchar(10),
	websites json,
	social_media json,
	logo varchar(150),
	created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

COMMIT;
