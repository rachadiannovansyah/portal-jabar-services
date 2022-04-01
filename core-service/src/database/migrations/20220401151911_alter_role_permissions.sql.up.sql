BEGIN;

ALTER TABLE roles 
  MODIFY COLUMN description varchar(100) NULL,
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

DROP TABLE IF EXISTS roles_permissions;
CREATE TABLE roles_permissions (
  role_id tinyint(2) unsigned NOT NULL,
  permission_id int(10) unsigned NOT NULL,
  PRIMARY KEY (role_id, permission_id),
  CONSTRAINT roles_permissions_role_id_fk FOREIGN KEY (role_id) REFERENCES roles (id),
  CONSTRAINT roles_permissions_permission_id_fk FOREIGN KEY (permission_id) REFERENCES permissions (id)
);

COMMIT;