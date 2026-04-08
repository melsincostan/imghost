-- Create tables
-- +migrate Up
create table galleries (
    internal_id serial primary key,
    slug blob not null,
    visibility int not null,
    metadata_privacy int not null,
    name text,
    description text,
    unique(slug)
);

create index idx_gallery_slug on galleries (slug);

create table images (
    internal_id serial primary key,
    slug blob not null,
    filename text not null,
    mimetype text not null,
    gallery_slug not null,
    unique(slug),
    constraint fk_gallery foreign key(gallery_slug) references galleries(slug) on delete cascade
);

create index idx_image_slug on images (slug); 


-- +migrate Down
drop index idx_image_slug;
drop table images;
drop index idx_gallery_slug;
drop table galleries;
