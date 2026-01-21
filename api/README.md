# Skill flow API

A golang api learning platform MVP.

- It implements simple authentication
- Course management
- Lesson management
- Enrollment system
- Progress tracking
- Uses file system for data storage instead of s3 or cloudinary for simplicity

## Data Scheme

1. Users - students/creators
2. Courses
3. Lessons
4. Enrollments
5. Progress - Progress % = (completed lessons/total lessons) * 100

## Endpoints

1. Auth
    - **POST** - `/auth/register`
    - **POST** - `/auth/login`
    - **GET** - `/auth/logout`

2. Courses
    - **GET** - `/courses` - list all courses
    - **GET** - `/courses/:id` - get course details
    - **POST** - `/courses/` (instructor only) - create a course
    - **PUT** - `/courses` (instructor only) - update a course
    - **DELETE** - `/courses/:id` (instructor only) - delete a course

3. Lessons
    - **GET** - `/courses/:courseId/lessons` - list lessons in a course
    - **GET** - `/courses/:courseId/lessons/:lessonId` - get lesson details
    - **POST** - `/courses/:courseId/lessons` (instructor only) - create a lesson
    - **PUT** - `/courses/:courseId/lessons/:lessonId` (instructor only) - update a lesson
    - **DELETE** - `/courses/:courseId/lessons/:lessonId` (instructor only) - delete a lesson

4. Enrollments
    - **POST** - `/courses/:courseId/enroll` - enroll in a course
    - **GET** - `/users/me/enrollments` - get user's enrollments

5. Progress
    - **GET** - `/progress/courses/:courseId` - get progress for a course
    - **POST** - `/progress/courses/:courseId/lessons/:lessonId` - mark lesson as completed

## Database Schema - mysql

```sql
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role ENUM('student', 'instructor') DEFAULT 'student',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

```sql
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
```

```sql
CREATE TABLE lessons (
    id INT AUTO_INCREMENT PRIMARY KEY,
    course_id INT,
    title VARCHAR(100) NOT NULL,
    description TEXT,
    content_path VARCHAR(255) NOT NULL,
    -- content_type ENUM('video', 'article', 'quiz') DEFAULT 'article',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE CASCADE
);
```

```sql
CREATE TABLE enrollments (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    course_id INT,
    enrolled_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (course_id) REFERENCES courses(id)
);
```

```sql
CREATE TABLE progress (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    course_id INT,
    lesson_id INT,
    is_completed BOOLEAN DEFAULT TRUE,
    completed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (course_id) REFERENCES courses(id),
    FOREIGN KEY (lesson_id) REFERENCES lessons(id)
);
```
