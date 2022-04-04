BEGIN;

ALTER TABLE roles 
  MODIFY COLUMN description varchar(100) NULL,
  ADD COLUMN level tinyint(2) NULL AFTER id,
  ADD COLUMN label varchar(100) NULL AFTER name;

CREATE UNIQUE INDEX idx_name ON roles (name);

DROP TABLE IF EXISTS permissions;
CREATE TABLE permissions (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(100) NOT NULL,
  description varchar(255),
  PRIMARY KEY (id)
);
CREATE UNIQUE INDEX idx_name ON permissions (name);

DROP TABLE IF EXISTS role_permissions;
CREATE TABLE role_permissions (
  role_id tinyint(2) unsigned NOT NULL,
  permission_id int(10) unsigned NOT NULL,
  PRIMARY KEY (role_id, permission_id)
);

COMMIT;