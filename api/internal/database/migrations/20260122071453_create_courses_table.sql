-- +goose Up
-- +goose StatementBegin
CREATE TABLE courses (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    thumbnail VARCHAR(255),
    price DECIMAL(10, 2) DEFAULT 0.00,
    is_published BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS courses;
-- +goose StatementEnd
