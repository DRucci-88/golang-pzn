-- Active: 1782153380096@@localhost@3306@golang_rest_pzn
CREATE TABLE category(
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(200) NOT NULL
) engine = InnoDB;

SELECT * FROM category;

DESC category;