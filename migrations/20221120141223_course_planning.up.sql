CREATE TABLE IF NOT EXISTS user_course_planning(
    user_id uuid,
    course_order INTEGER
    course_name VARCHAR(255) NOT NULL,
    PRIMARY KEY (user_id, course_order)
);