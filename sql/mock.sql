-- Profiles
INSERT INTO PROFILE (username, last_seen, date_joined) VALUES ('Soul', '12-12-12', '12-12-12');
INSERT INTO PROFILE (username, last_seen, date_joined) VALUES ('Ty', '12-12-13', '12-12-13');
INSERT INTO PROFILE (username, last_seen, date_joined) VALUES ('Ankit', '12-12-13', '12-12-13');
INSERT INTO PROFILE (username, last_seen, date_joined) VALUES ('Abdullah', '12-12-14', '12-12-14');

-- Contents
INSERT INTO CONTENT (title, summary, description, last_modified, date_created)
    VALUES ('Title A', 'This is a summary', 'This is a description', '12-12-12', '12-12-12');
INSERT INTO CONTENT (title, summary, description, last_modified, date_created)
    VALUES ('Title B', 'This is a summary 2', 'This is a description 2', '12-12-12', '12-12-12');
INSERT INTO CONTENT (title, summary, description, last_modified, date_created)
    VALUES ('Title C', 'This is a summary 3', 'This is a description 3', '12-12-12', '12-12-12');
INSERT INTO CONTENT (title, summary, description, last_modified, date_created)
    VALUES ('Title D', 'This is a summary 4', 'This is a description 4', '12-12-12', '12-12-12');

-- Content Author

INSERT INTO CONTENT_PROFILES (id, username) VALUES (1, 'Soul');
INSERT INTO CONTENT_PROFILES (id, username) VALUES (1, 'Ty');
INSERT INTO CONTENT_PROFILES (id, username) VALUES (2, 'Ty');
INSERT INTO CONTENT_PROFILES (id, username) VALUES (3, 'Soul');
INSERT INTO CONTENT_PROFILES (id, username) VALUES (3, 'Ty');
INSERT INTO CONTENT_PROFILES (id, username) VALUES (3, 'Ankit');
INSERT INTO CONTENT_PROFILES (id, username) VALUES (3, 'Abdullah');
INSERT INTO CONTENT_PROFILES (id, username) VALUES (4, 'Ankit');
INSERT INTO CONTENT_PROFILES (id, username) VALUES (4, 'Abdullah');

-- Tags
INSERT INTO TAG (tag) VALUES ('YOLO');
INSERT INTO TAG (tag) VALUES ('SWAG');
INSERT INTO TAG (tag) VALUES ('AgilexInterns');
INSERT INTO TAG (tag) VALUES ('Sushi');

-- Content Tags
INSERT INTO CONTENT_TAG (id, tag) VALUES (1, 'YOLO');
INSERT INTO CONTENT_TAG (id, tag) VALUES (1, 'SWAG');
INSERT INTO CONTENT_TAG (id, tag) VALUES (1, 'AgilexInterns');
INSERT INTO CONTENT_TAG (id, tag) VALUES (1, 'Sushi');
INSERT INTO CONTENT_TAG (id, tag) VALUES (2, 'YOLO');
INSERT INTO CONTENT_TAG (id, tag) VALUES (3, 'YOLO');
INSERT INTO CONTENT_TAG (id, tag) VALUES (2, 'AgilexInterns');
INSERT INTO CONTENT_TAG (id, tag) VALUES (3, 'AgilexInterns');

-- Comments
INSERT INTO COMMENT (username, id, post_number, comment, last_modified, date_created)
    VALUES ('Soul', 1, 1, 'This is a comment', '12-12-12', '12-12-12');
INSERT INTO COMMENT (username, id, post_number, comment, last_modified, date_created)
    VALUES ('Soul', 2, 1, 'This is a comment 2', '12-12-12', '12-12-12');
INSERT INTO COMMENT (username, id, post_number, comment, last_modified, date_created)
    VALUES ('Soul', 1, 2, 'This is a comment 3', '12-12-12', '12-12-12');

-- Charges
INSERT INTO CHARGE (charged) VALUES (TRUE);
INSERT INTO CHARGE (charged) VALUES (TRUE);
INSERT INTO CHARGE (charged) VALUES (TRUE);
INSERT INTO CHARGE (charged) VALUES (TRUE);
INSERT INTO CHARGE (charged) VALUES (TRUE);
INSERT INTO CHARGE (charged) VALUES (TRUE);

-- Profile charges on Content
INSERT INTO PROFILE_CHARGES (id, username, charge_id) VALUES (1, 'Soul', 1);
INSERT INTO PROFILE_CHARGES (id, username, charge_id) VALUES (1, 'Ty', 2);
INSERT INTO PROFILE_CHARGES (id, username, charge_id) VALUES (1, 'Ankit', 3);
INSERT INTO PROFILE_CHARGES (id, username, charge_id) VALUES (2, 'Soul', 4);
INSERT INTO PROFILE_CHARGES (id, username, charge_id) VALUES (3, 'Abdullah', 5);
INSERT INTO PROFILE_CHARGES (id, username, charge_id) VALUES (3, 'Ankit', 6);

UPDATE CONTENT SET charges = charges + 1 WHERE id = 1;
UPDATE CONTENT SET charges = charges + 1 WHERE id = 1;
UPDATE CONTENT SET charges = charges + 1 WHERE id = 1;
UPDATE CONTENT SET charges = charges + 1 WHERE id = 2;
UPDATE CONTENT SET charges = charges + 1 WHERE id = 3;
UPDATE CONTENT SET charges = charges + 1 WHERE id = 3;