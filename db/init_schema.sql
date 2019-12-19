use astral;

CREATE TABLE users (
    id MEDIUMINT NOT NULL AUTO_INCREMENT,
    login CHAR(30) NOT NULL,
    password CHAR(30) NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO users (login, password) VALUES
('test', '123'), ('alex', 'qwerty');


CREATE TABLE tokens (
    id MEDIUMINT NOT NULL AUTO_INCREMENT,
    ownerid MEDIUMINT NOT NULL,
    token CHAR(30) NOT NULL,
    expiration_date DATETIME NULL,
    PRIMARY KEY (id)
);

CREATE TABLE links (
    id MEDIUMINT NOT NULL AUTO_INCREMENT,
    ownerid MEDIUMINT NOT NULL,
    short CHAR(30) NOT NULL,
    full VARCHAR(255) NOT NULL,
    expiration_date DATETIME NULL,
    PRIMARY KEY (id)
);

CREATE TABLE statistics (
    id MEDIUMINT NOT NULL AUTO_INCREMENT,
    short CHAR(30) NOT NULL,
    full VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);
