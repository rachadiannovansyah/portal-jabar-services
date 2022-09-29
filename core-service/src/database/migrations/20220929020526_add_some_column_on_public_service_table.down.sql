ALTER TABLE public_services
DROP COLUMN excerpt,
DROP COLUMN slug,
DROP COLUMN service_type,
DROP COLUMN websites,
DROP COLUMN social_media,
DROP COLUMN video,
DROP COLUMN purposes,
DROP COLUMN facilities,
DROP COLUMN info;
ALTER TABLE public_services CHANGE images image varchar(100) DEFAULT NULL;
DROP INDEX idx_name ON public_services;