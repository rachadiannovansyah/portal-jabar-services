ALTER TABLE public_services
ADD excerpt varchar(150) NULL AFTER description,
ADD slug varchar(100) UNIQUE NULL AFTER excerpt,
ADD service_type varchar(30) NULL AFTER slug,
ADD websites json NULL AFTER service_type,
ADD social_media json NULL AFTER websites,
ADD video varchar(100) NULL AFTER social_media,
ADD purposes json NULL AFTER video,
ADD facilities json NULL after purposes,
ADD info json NULL after facilities;

CREATE INDEX idx_slug ON public_services (slug);
CREATE INDEX idx_name ON public_services (name);

ALTER TABLE public_services CHANGE image images json NULL;
