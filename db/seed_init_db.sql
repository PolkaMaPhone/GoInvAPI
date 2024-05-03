-- noinspection SqlResolveForFile

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

DROP SCHEMA IF EXISTS test_seed CASCADE;
CREATE SCHEMA test_seed;
ALTER SCHEMA test_seed OWNER TO postgres;
SET default_tablespace = '';
SET default_table_access_method = heap;

CREATE TABLE test_seed.categories
(
    category_id integer                  NOT NULL,
    name        character varying(255)   NOT NULL,
    description text,
    "createdAt" timestamp with time zone NOT NULL,
    "updatedAt" timestamp with time zone NOT NULL
);

ALTER TABLE test_seed.categories
    OWNER TO postgres;

CREATE SEQUENCE test_seed.categories_category_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE test_seed.categories_category_id_seq OWNER TO postgres;

ALTER SEQUENCE test_seed.categories_category_id_seq OWNED BY test_seed.categories.category_id;

CREATE TABLE test_seed.groups
(
    group_id    integer                  NOT NULL,
    name        character varying(255)   NOT NULL,
    description text,
    "createdAt" timestamp with time zone NOT NULL,
    "updatedAt" timestamp with time zone NOT NULL
);

ALTER TABLE test_seed.groups
    OWNER TO postgres;

CREATE SEQUENCE test_seed.groups_group_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE test_seed.groups_group_id_seq OWNER TO postgres;

ALTER SEQUENCE test_seed.groups_group_id_seq OWNED BY test_seed.groups.group_id;


CREATE TABLE test_seed.item_history
(
    history_id  integer,
    item_id     integer,
    date        date,
    notes       text,
    "createdAt" timestamp with time zone NOT NULL,
    "updatedAt" timestamp with time zone NOT NULL
);

ALTER TABLE test_seed.item_history
    OWNER TO postgres;

CREATE TABLE test_seed.items
(
    item_id     integer                  NOT NULL,
    name        character varying(255)   NOT NULL,
    description text,
    category_id integer,
    group_id    integer,
    location_id integer,
    is_stored   boolean DEFAULT false,
    "createdAt" timestamp with time zone NOT NULL,
    "updatedAt" timestamp with time zone NOT NULL
);


ALTER TABLE test_seed.items
    OWNER TO postgres;

CREATE SEQUENCE test_seed.items_item_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE test_seed.items_item_id_seq OWNER TO postgres;

ALTER SEQUENCE test_seed.items_item_id_seq OWNED BY test_seed.items.item_id;

CREATE TABLE test_seed.locations
(
    location_id integer                  NOT NULL,
    tub_id      integer,
    shelf_id    integer,
    "createdAt" timestamp with time zone NOT NULL,
    "updatedAt" timestamp with time zone NOT NULL
);

ALTER TABLE test_seed.locations
    OWNER TO postgres;

CREATE SEQUENCE test_seed.locations_location_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE test_seed.locations_location_id_seq OWNER TO postgres;

ALTER SEQUENCE test_seed.locations_location_id_seq OWNED BY test_seed.locations.location_id;

CREATE TABLE test_seed.shelves
(
    shelf_id    integer                  NOT NULL,
    label       character varying(255)   NOT NULL,
    description text,
    location_id integer,
    "createdAt" timestamp with time zone NOT NULL,
    "updatedAt" timestamp with time zone NOT NULL
);

ALTER TABLE test_seed.shelves
    OWNER TO postgres;

CREATE SEQUENCE test_seed.shelves_shelf_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE test_seed.shelves_shelf_id_seq OWNER TO postgres;

ALTER SEQUENCE test_seed.shelves_shelf_id_seq OWNED BY test_seed.shelves.shelf_id;

CREATE TABLE test_seed.transactions
(
    transaction_id integer,
    item_id        integer,
    type           character varying(50),
    "timestamp"    timestamp without time zone,
    user_id        integer,
    "createdAt"    timestamp with time zone NOT NULL,
    "updatedAt"    timestamp with time zone NOT NULL
);

ALTER TABLE test_seed.transactions
    OWNER TO postgres;

CREATE TABLE test_seed.tubs
(
    tub_id      integer                  NOT NULL,
    label       character varying(255)   NOT NULL,
    shelf_id    integer                  NOT NULL,
    location_id integer,
    "createdAt" timestamp with time zone NOT NULL,
    "updatedAt" timestamp with time zone NOT NULL
);


ALTER TABLE test_seed.tubs
    OWNER TO postgres;

CREATE SEQUENCE test_seed.tubs_tub_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE test_seed.tubs_tub_id_seq OWNER TO postgres;

ALTER SEQUENCE test_seed.tubs_tub_id_seq OWNED BY test_seed.tubs.tub_id;

ALTER TABLE ONLY test_seed.categories
    ALTER COLUMN category_id SET DEFAULT nextval('test_seed.categories_category_id_seq'::regclass);

ALTER TABLE ONLY test_seed.groups
    ALTER COLUMN group_id SET DEFAULT nextval('test_seed.groups_group_id_seq'::regclass);

ALTER TABLE ONLY test_seed.items
    ALTER COLUMN item_id SET DEFAULT nextval('test_seed.items_item_id_seq'::regclass);

ALTER TABLE ONLY test_seed.locations
    ALTER COLUMN location_id SET DEFAULT nextval('test_seed.locations_location_id_seq'::regclass);

ALTER TABLE ONLY test_seed.shelves
    ALTER COLUMN shelf_id SET DEFAULT nextval('test_seed.shelves_shelf_id_seq'::regclass);

ALTER TABLE ONLY test_seed.tubs
    ALTER COLUMN tub_id SET DEFAULT nextval('test_seed.tubs_tub_id_seq'::regclass);