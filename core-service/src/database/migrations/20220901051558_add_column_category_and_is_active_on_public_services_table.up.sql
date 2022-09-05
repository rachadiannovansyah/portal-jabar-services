ALTER TABLE public_services 
ADD COLUMN category varchar(100) NULL AFTER image,
ADD COLUMN is_active tinyint(1) NULL after category;