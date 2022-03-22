BEGIN;
CREATE TABLE registration_invitations (
  id INT(10) unsigned NOT NULL AUTO_INCREMENT,
  email VARCHAR(80) NOT NULL,
  token VARCHAR(255),
  expired_at TIMESTAMP NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
CREATE INDEX idx_email ON registration_invitations(email);
CREATE INDEX idx_token ON registration_invitations(token);
COMMIT;