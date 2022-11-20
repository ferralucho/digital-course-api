CREATE TABLE IF NOT EXISTS user_course_planning(
    user_id uuid NOT NULL,
    course_order INTEGER NOT NULL,
    course_name VARCHAR(255) NOT NULL,
    CONSTRAINT user_course_planning_pkey PRIMARY KEY (user_id, course_order)
);