CREATE TABLE news
(
    id        serial       not null unique,
    category  varchar(255) not null,
    datetime  int          not null,
    headline  varchar(255) not null,
    image     varchar(255) not null,
    related   varchar(255) not null,
    resource  varchar(255) not null,
    summary   varchar(512) not null,
    url       varchar(255) not null
);