ALTER TABLE users
  ADD COLUMN nip varchar(18),
  ADD COLUMN occupation varchar(35)
  AFTER `password`;