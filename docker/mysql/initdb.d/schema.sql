CREATE TABLE users (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(32) NOT NULL,
    email VARCHAR(32) NOT NULL,
    pass_hash VARCHAR(256) NOT NULL,
    PRIMARY KEY (id)
)DEFAULT CHARACTER SET=utf8mb4;