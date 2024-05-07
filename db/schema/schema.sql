create table item_history
(
    history_id  integer,
    item_id     integer,
    date        date,
    notes       text,
    "createdAt" timestamp with time zone not null default NOW(),
    "updatedAt" timestamp with time zone not null default NOW()
);

alter table item_history
    owner to postgres;

create table transactions
(
    transaction_id integer,
    item_id        integer,
    type           varchar(50),
    timestamp      timestamp,
    user_id        integer,
    "createdAt" timestamp with time zone not null default NOW(),
    "updatedAt" timestamp with time zone not null default NOW()
);

alter table transactions
    owner to postgres;

create table categories
(
    category_id serial
        primary key,
    name        varchar(255)             not null,
    description text,
    "createdAt" timestamp with time zone not null default NOW(),
    "updatedAt" timestamp with time zone not null default NOW()
);

alter table categories
    owner to postgres;

create table groups
(
    group_id    serial
        primary key,
    name        varchar(255)             not null,
    description text,
    "createdAt" timestamp with time zone not null default NOW(),
    "updatedAt" timestamp with time zone not null default NOW()
);

alter table groups
    owner to postgres;

create table locations
(
    location_id serial
        primary key,
    tub_id      integer,
    shelf_id    integer,
    "createdAt" timestamp with time zone not null default NOW(),
    "updatedAt" timestamp with time zone not null default NOW()
);

alter table locations
    owner to postgres;

create table items
(
    item_id     serial
        primary key,
    name        varchar(255)             not null,
    description text,
    category_id integer
        references categories
            on update cascade on delete cascade,
    group_id    integer
        references groups
            on update cascade on delete cascade,
    location_id integer
        references locations
            on update cascade on delete cascade,
    is_stored   boolean default false,
    "createdAt" timestamp with time zone not null default NOW(),
    "updatedAt" timestamp with time zone not null default NOW()
);

alter table items
    owner to postgres;

create table shelves
(
    shelf_id    serial
        primary key,
    label       varchar(255)             not null,
    description text,
    location_id integer
                                         references locations
                                             on update cascade on delete set null,
    "createdAt" timestamp with time zone not null default NOW(),
    "updatedAt" timestamp with time zone not null default NOW()
);

alter table shelves
    owner to postgres;

create table tubs
(
    tub_id      serial
        primary key,
    label       varchar(255)             not null,
    shelf_id    integer                  not null
        references shelves
            on update cascade on delete cascade,
    location_id integer
        references locations
            on update cascade on delete cascade,
    "createdAt" timestamp with time zone not null default NOW(),
    "updatedAt" timestamp with time zone not null default NOW()
);

alter table tubs
    owner to postgres;

