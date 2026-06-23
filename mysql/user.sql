CREATE TABLE user(
    username VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    PRIMARY KEY (username)
);

SELECT * FROM `user`

INSERT INTO user(username, password) VALUES ('admin', 'admin')