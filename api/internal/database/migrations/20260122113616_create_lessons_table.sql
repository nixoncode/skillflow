-- +goose Up
-- +goose StatementBegin
CREATE TABLE lessons (
    id INT AUTO_INCREMENT PRIMARY KEY,
    course_id INT,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    content_path VARCHAR(255) NOT NULL,
    -- content_type ENUM('video', 'article', 'quiz') DEFAULT 'article',
    -- position INT NOT NULL,  -- To keep lessons in order (1, 2, 3...)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS lessons;
-- +goose StatementEnd
