-- noinspection SqlResolveForFile

-- noinspection SpellCheckingInspectionForFile

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

INSERT INTO test_seed.categories (category_id, name, description, "createdAt", "updatedAt")
VALUES (1, 'Sample Category 1', 'Description for Sample Category 1', NOW(),
        NOW());
INSERT INTO test_seed.categories (category_id, name, description, "createdAt", "updatedAt")
VALUES (2, 'Sample Category 2', 'Description for Sample Category 2', NOW(),
        NOW());

INSERT INTO test_seed.groups (group_id, name, description, "createdAt", "updatedAt")
VALUES (1, 'Sample Group 1', 'Description for Sample Group 1', NOW(),
        NOW());
INSERT INTO test_seed.groups (group_id, name, description, "createdAt", "updatedAt")
VALUES (2, 'Sample Group 2', 'Description for Sample Group 2', NOW(),
        NOW());
INSERT INTO test_seed.groups (group_id, name, description, "createdAt", "updatedAt")
VALUES (3, 'Ungrouped', 'This item does not belong to any specific group.', NOW(),
        NOW());

INSERT INTO test_seed.items (item_id, name, description, category_id, group_id, location_id, is_stored, "createdAt",
                             "updatedAt")
VALUES (1, 'Sample Item 111', 'Description for Sample Item 1-1-1', 1, 1, 1, false, NOW(),
        NOW());
INSERT INTO test_seed.items (item_id, name, description, category_id, group_id, location_id, is_stored, "createdAt",
                             "updatedAt")
VALUES (2, 'Sample Item 112', 'Description for Sample Item 1-1-2', 1, 1, 2, false, NOW(),
        NOW());
INSERT INTO test_seed.items (item_id, name, description, category_id, group_id, location_id, is_stored, "createdAt",
                             "updatedAt")
VALUES (3, 'Sample Item 113', 'Description for Sample Item 1-1-3', 1, 1, 3, false, NOW(),
        NOW());
INSERT INTO test_seed.items (item_id, name, description, category_id, group_id, location_id, is_stored, "createdAt",
                             "updatedAt")
VALUES (4, 'Sample Item 114', 'Description for Sample Item 1-1-4', 1, 1, 4, false, NOW(),
        NOW());
INSERT INTO test_seed.items (item_id, name, description, category_id, group_id, location_id, is_stored, "createdAt",
                             "updatedAt")
VALUES (5, 'Sample Item 121', 'Description for Sample Item 1-2-1', 1, 2, 1, false, NOW(),
        NOW());
INSERT INTO test_seed.items (item_id, name, description, category_id, group_id, location_id, is_stored, "createdAt",
                             "updatedAt")
VALUES (6, 'Sample Item 122', 'Description for Sample Item 1-2-2', 1, 2, 2, false, NOW(),
        NOW());
INSERT INTO test_seed.items (item_id, name, description, category_id, group_id, location_id, is_stored, "createdAt",
                             "updatedAt")
VALUES (7, 'Sample Item 123', 'Description for Sample Item 1-2-3', 1, 2, 3, false, NOW(),
        NOW());
INSERT INTO test_seed.items (item_id, name, description, category_id, group_id, location_id, is_stored, "createdAt",
                             "updatedAt")
VALUES (8, 'Sample Item 124', 'Description for Sample Item 1-2-4', 1, 2, 4, false, NOW(),
        NOW());
INSERT INTO test_seed.items (item_id, name, description, category_id, group_id, location_id, is_stored, "createdAt",
                             "updatedAt")
VALUES (9, 'Sample Item 131', 'Description for Sample Item 1-3-1', 1, 3, 1, false, NOW(),
        NOW());
INSERT INTO test_seed.items (item_id, name, description, category_id, group_id, location_id, is_stored, "createdAt",
                             "updatedAt")
VALUES (10, 'Sample Item 132', 'Description for Sample Item 1-3-2', 1, 3, 2, false, NOW(),
        NOW());
INSERT INTO test_seed.items (item_id, name, description, category_id, group_id, location_id, is_stored, "createdAt",
                             "updatedAt")
VALUES (11, 'Sample Item 133', 'Description for Sample Item 1-3-3', 1, 3, 3, false, NOW(),
        NOW());
INSERT INTO test_seed.items (item_id, name, description, category_id, group_id, location_id, is_stored, "createdAt",
                             "updatedAt")
VALUES (12, 'Sample Item 134', 'Description for Sample Item 1-3-4', 1, 3, 4, false, NOW(),
        NOW());
INSERT INTO test_seed.items (item_id, name, description, category_id, group_id, location_id, is_stored, "createdAt",
                             "updatedAt")
VALUES (13, 'Sample Item 211', 'Description for Sample Item 2-1-1', 2, 1, 1, false, NOW(),
        NOW());
INSERT INTO test_seed.items (item_id, name, description, category_id, group_id, location_id, is_stored, "createdAt",
                             "updatedAt")
VALUES (14, 'Sample Item 212', 'Description for Sample Item 2-1-2', 2, 1, 2, false, NOW(),
        NOW());
INSERT INTO test_seed.items (item_id, name, description, category_id, group_id, location_id, is_stored, "createdAt",
                             "updatedAt")
VALUES (15, 'Sample Item 213', 'Description for Sample Item 2-1-3', 2, 1, 3, false, NOW(),
        NOW());
INSERT INTO test_seed.items (item_id, name, description, category_id, group_id, location_id, is_stored, "createdAt",
                             "updatedAt")
VALUES (16, 'Sample Item 214', 'Description for Sample Item 2-1-4', 2, 1, 4, false, NOW(),
        NOW());
INSERT INTO test_seed.items (item_id, name, description, category_id, group_id, location_id, is_stored, "createdAt",
                             "updatedAt")
VALUES (17, 'Sample Item 221', 'Description for Sample Item 2-2-1', 2, 2, 1, false, NOW(),
        NOW());
INSERT INTO test_seed.items (item_id, name, description, category_id, group_id, location_id, is_stored, "createdAt",
                             "updatedAt")
VALUES (18, 'Sample Item 222', 'Description for Sample Item 2-2-2', 2, 2, 2, false, NOW(),
        NOW());
INSERT INTO test_seed.items (item_id, name, description, category_id, group_id, location_id, is_stored, "createdAt",
                             "updatedAt")
VALUES (19, 'Sample Item 223', 'Description for Sample Item 2-2-3', 2, 2, 3, false, NOW(),
        NOW());
INSERT INTO test_seed.items (item_id, name, description, category_id, group_id, location_id, is_stored, "createdAt",
                             "updatedAt")
VALUES (20, 'Sample Item 224', 'Description for Sample Item 2-2-4', 2, 2, 4, false, NOW(),
        NOW());
INSERT INTO test_seed.items (item_id, name, description, category_id, group_id, location_id, is_stored, "createdAt",
                             "updatedAt")
VALUES (21, 'Sample Item 231', 'Description for Sample Item 2-3-1', 2, 3, 1, false, NOW(),
        NOW());
INSERT INTO test_seed.items (item_id, name, description, category_id, group_id, location_id, is_stored, "createdAt",
                             "updatedAt")
VALUES (22, 'Sample Item 232', 'Description for Sample Item 2-3-2', 2, 3, 2, false, NOW(),
        NOW());
INSERT INTO test_seed.items (item_id, name, description, category_id, group_id, location_id, is_stored, "createdAt",
                             "updatedAt")
VALUES (23, 'Sample Item 233', 'Description for Sample Item 2-3-3', 2, 3, 3, false, NOW(),
        NOW());
INSERT INTO test_seed.items (item_id, name, description, category_id, group_id, location_id, is_stored, "createdAt",
                             "updatedAt")
VALUES (24, 'Sample Item 234', 'Description for Sample Item 2-3-4', 2, 3, 4, false, NOW(),
        NOW());
INSERT INTO test_seed.items (item_id, name, description, category_id, group_id, location_id, is_stored, "createdAt",
                             "updatedAt")
VALUES (25, 'Uncategorized Item 1', 'Description for Uncategorized Item 1', NULL, NULL, NULL, false,
        NOW(), NOW());
INSERT INTO test_seed.items (item_id, name, description, category_id, group_id, location_id, is_stored, "createdAt",
                             "updatedAt")
VALUES (26, 'Uncategorized Item 2', 'Description for Uncategorized Item 2', NULL, NULL, NULL, false,
        NOW(), NOW());
INSERT INTO test_seed.items (item_id, name, description, category_id, group_id, location_id, is_stored, "createdAt",
                             "updatedAt")
VALUES (27, 'Uncategorized Item 3', 'Description for Uncategorized Item 3', NULL, NULL, NULL, false,
        NOW(), NOW());

INSERT INTO test_seed.locations (location_id, tub_id, shelf_id, "createdAt", "updatedAt")
VALUES (1, NULL, 1, NOW(), NOW());
INSERT INTO test_seed.locations (location_id, tub_id, shelf_id, "createdAt", "updatedAt")
VALUES (2, NULL, 2, NOW(), NOW());
INSERT INTO test_seed.locations (location_id, tub_id, shelf_id, "createdAt", "updatedAt")
VALUES (3, 1, 1, NOW(), NOW());
INSERT INTO test_seed.locations (location_id, tub_id, shelf_id, "createdAt", "updatedAt")
VALUES (4, 2, 2, NOW(), NOW());
INSERT INTO test_seed.locations (location_id, tub_id, shelf_id, "createdAt", "updatedAt")
VALUES (0, 0, 0, NOW(), NOW());

INSERT INTO test_seed.shelves (shelf_id, label, description, location_id, "createdAt", "updatedAt")
VALUES (1, 'Shelf 1', NULL, 1, NOW(), NOW());
INSERT INTO test_seed.shelves (shelf_id, label, description, location_id, "createdAt", "updatedAt")
VALUES (2, 'Shelf 2', NULL, 2, NOW(), NOW());
INSERT INTO test_seed.shelves (shelf_id, label, description, location_id, "createdAt", "updatedAt")
VALUES (3, 'Uncategorized', NULL, 0, NOW(), NOW());

INSERT INTO test_seed.tubs (tub_id, label, shelf_id, location_id, "createdAt", "updatedAt")
VALUES (1, 'Tub 1', 1, 3, NOW(), NOW());
INSERT INTO test_seed.tubs (tub_id, label, shelf_id, location_id, "createdAt", "updatedAt")
VALUES (2, 'Tub 2', 2, 4, NOW(), NOW());

SELECT pg_catalog.setval('test_seed.categories_category_id_seq', 2, true);

SELECT pg_catalog.setval('test_seed.groups_group_id_seq', 3, true);

SELECT pg_catalog.setval('test_seed.items_item_id_seq', 27, true);

SELECT pg_catalog.setval('test_seed.locations_location_id_seq', 1, false);

SELECT pg_catalog.setval('test_seed.shelves_shelf_id_seq', 3, true);

SELECT pg_catalog.setval('test_seed.tubs_tub_id_seq', 2, true);

ALTER TABLE ONLY test_seed.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (category_id);

ALTER TABLE ONLY test_seed.groups
    ADD CONSTRAINT groups_pkey PRIMARY KEY (group_id);

ALTER TABLE ONLY test_seed.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (item_id);

ALTER TABLE ONLY test_seed.locations
    ADD CONSTRAINT locations_pkey PRIMARY KEY (location_id);

ALTER TABLE ONLY test_seed.shelves
    ADD CONSTRAINT shelves_pkey PRIMARY KEY (shelf_id);

ALTER TABLE ONLY test_seed.tubs
    ADD CONSTRAINT tubs_pkey PRIMARY KEY (tub_id);

ALTER TABLE ONLY test_seed.items
    ADD CONSTRAINT items_category_id_fkey FOREIGN KEY (category_id) REFERENCES test_seed.categories (category_id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY test_seed.items
    ADD CONSTRAINT items_group_id_fkey FOREIGN KEY (group_id) REFERENCES test_seed.groups (group_id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY test_seed.items
    ADD CONSTRAINT items_location_id_fkey FOREIGN KEY (location_id) REFERENCES test_seed.locations (location_id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY test_seed.shelves
    ADD CONSTRAINT shelves_location_id_fkey FOREIGN KEY (location_id) REFERENCES test_seed.locations (location_id) ON UPDATE CASCADE ON DELETE SET NULL;

ALTER TABLE ONLY test_seed.tubs
    ADD CONSTRAINT tubs_location_id_fkey FOREIGN KEY (location_id) REFERENCES test_seed.locations (location_id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE ONLY test_seed.tubs
    ADD CONSTRAINT tubs_shelf_id_fkey FOREIGN KEY (shelf_id) REFERENCES test_seed.shelves (shelf_id) ON UPDATE CASCADE ON DELETE CASCADE;