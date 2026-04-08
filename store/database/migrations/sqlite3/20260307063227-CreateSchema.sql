-- Create tables
-- +migrate Up
create table galleries (            -- a gallery is a collection of images
    slug blob not null primary key, -- Should be a set number of random bytes
    visibility int not null,        -- int denoting visibility level                            (should default to the most private setting)
    metadata_privacy int not null,  -- int denoting whether to strip metadata / only some / all (should default to all)
    name text,                      -- name of the gallery (optional)
    description text                -- description of the gallery (optional)
);

create table images (               -- an image stores information about image files
    slug blob not null primary key, -- should be a set number of random bytes
    filename text not null,         -- original filename of the image file
    mimetype text not null,         -- mime type of the original image
    gallery_slug not null           -- sqlite needs foreign keys to be explicitely enabled, so not using them here out of convenience
);

create index idx_image_gallery on images (
    gallery_slug
);


-- +migrate Down
drop index idx_image_gallery;
drop table images;
drop table galleries;
