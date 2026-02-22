CREATE TABLE users(
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL DEFAULT '',
    age INT NOT NULL DEFAULT 0,
    phone VARCHAR(50) NOT NULL DEFAULT ''
);

INSERT INTO users (name, email, age, phone) VALUES
('John Doe', 'john@example.com', 30, '+12345678901');