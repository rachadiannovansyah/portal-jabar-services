ALTER TABLE pop_up_banners 
ADD COLUMN is_live tinyint(1)
DEFAULT 0
AFTER duration;