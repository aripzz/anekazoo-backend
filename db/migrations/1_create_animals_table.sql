-- migrate:up
CREATE TABLE IF NOT EXISTS animals (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    class VARCHAR(255) NOT NULL,
    legs INT NOT NULL
);

-- migrate:down
DROP TABLE IF EXISTS animals;

