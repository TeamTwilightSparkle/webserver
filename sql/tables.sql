CREATE TABLE IF NOT EXISTS PROFILE (
    id  SERIAL PRIMARY KEY,
    username    VARCHAR(32) NOT NULL UNIQUE,
    level       INT DEFAULT 0 CHECK (level >= 0), -- Administrative level
    last_seen   DATE        NOT NULL,
    date_joined DATE        NOT NULL
);

CREATE TABLE IF NOT EXISTS CONTENT (
    id            SERIAL PRIMARY KEY,
    title         VARCHAR(64)  NOT NULL,
    summary       VARCHAR(160) NOT NULL,
    description   VARCHAR(512) NOT NULL,
    image         VARCHAR(255),
    charges       INT DEFAULT 0 CHECK (charges >= 0),
    last_modified DATE         NOT NULL,
    date_created  DATE         NOT NULL
);

CREATE TABLE IF NOT EXISTS TAG (
    tag_id  SERIAL PRIMARY KEY,
    tag VARCHAR(32) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS COMMENT (
    comment_id    SERIAL PRIMARY KEY,
    username      VARCHAR(32) REFERENCES PROFILE (username),
    id            INT REFERENCES CONTENT (id),
    post_number   INT DEFAULT 1,
    comment       VARCHAR(160),
    last_modified DATE NOT NULL,
    date_created  DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS CHARGE (
    charge_id      SERIAL PRIMARY KEY,
    charged BOOLEAN DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS CONTENT_PROFILES (
    id         INT REFERENCES CONTENT (id),
    username   VARCHAR(32) REFERENCES PROFILE (username),
    PRIMARY KEY (id, username)
);

CREATE TABLE IF NOT EXISTS PROFILE_CHARGES (
    id         INT REFERENCES CONTENT (id),
    username   VARCHAR(32) REFERENCES PROFILE (username),
    charge_id  INT REFERENCES CHARGE (charge_id),
    PRIMARY KEY (username, charge_id, id)
);

CREATE TABLE IF NOT EXISTS CONTENT_TAG (
    id         INT REFERENCES CONTENT (id),
    tag        VARCHAR(32) REFERENCES TAG (tag),
    PRIMARY KEY (id, tag)
);