
CREATE TABLE comments
(
    id INT NOT NULL AUTO_INCREMENT,
    email VARCHAR(100) UNIQUE NOT NULL,
    comment TEXT,
    PRIMARY KEY (id)
);

DROP TABLE comments

DESC comments;

SELECT * FROM comments;