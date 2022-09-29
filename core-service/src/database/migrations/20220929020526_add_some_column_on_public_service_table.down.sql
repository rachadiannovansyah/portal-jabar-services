ALTER TABLE public_services
DROP excerpt,
DROP slug,
DROP service_type,
DROP website,
DROP social_media,
DROP video,
DROP purposes,
DROP facilities,
DROP info,
DROP logo;
ALTER TABLE public_services CHANGE images image varchar(100) DEFAULT NULL;
DROP INDEX idx_name ON public_services;
