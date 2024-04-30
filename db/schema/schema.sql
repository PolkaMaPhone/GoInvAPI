create table item_history
(
    history_id  integer,
    item_id     integer,
    date        date,
    notes       text,
    "createdAt" timestamp with time zone not null,
    "updatedAt" timestamp with time zone not null
);

create table transactions
(
    transaction_id integer,
    item_id        integer,
    type           varchar(50),
    timestamp      timestamp,
    user_id        integer,
    "createdAt"    timestamp with time zone not null,
    "updatedAt"    timestamp with time zone not null
);

create table categories
(
    category_id serial
        primary key,
    name        varchar(255)             not null,
    description text,
    "createdAt" timestamp with time zone not null,
    "updatedAt" timestamp with time zone not null
);

create table groups
(
    group_id    serial
        primary key,
    name        varchar(255)             not null,
    description text,
    "createdAt" timestamp with time zone not null,
    "updatedAt" timestamp with time zone not null
);

create table locations
(
    location_id serial
        primary key,
    tub_id      integer,
    shelf_id    integer,
    "createdAt" timestamp with time zone not null,
    "updatedAt" timestamp with time zone not null
);

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
    "createdAt" timestamp with time zone not null,
    "updatedAt" timestamp with time zone not null
);

create table shelves
(
    shelf_id    serial
        primary key,
    label       varchar(255)             not null,
    description text,
    location_id integer
                                         references locations
                                             on update cascade on delete set null,
    "createdAt" timestamp with time zone not null,
    "updatedAt" timestamp with time zone not null
);

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
    "createdAt" timestamp with time zone not null,
    "updatedAt" timestamp with time zone not null
);

