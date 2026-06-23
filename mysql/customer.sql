CREATE TABLE customer
(
    id VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL,
    PRIMARY KEY (id)
);

ALTER TABLE customer
    ADD COLUMN email VARCHAR(100),
    ADD COLUMN balance INTEGER DEFAULT 0,
    ADD COLUMN  rating DOUBLE DEFAULT 0.0,
    ADD COLUMN  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ADD COLUMN birth_date DATE,
    ADD COLUMN married BOOLEAN DEFAULT FALSE
;

INSERT INTO customer(id, name, email, balance, rating, birth_date, married)
VALUES
('rucci', "Rucci", "rucci@gmail.com", 1000000, 90.0, '2000-01-01', true),
('le', 'Le', 'le@gmail.com', 500000, 90.0, '2000-02-02', true)
;

UPDATE customer
SET email = NULL, birth_date = NULL
WHERE id = 'le'
;

DESC customer;

SELECT * FROM customer

SHOW TABLES;

DROP TABLE customer;

DELETE FROM customer;