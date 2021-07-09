BEGIN;

CREATE TABLE students (
    id varchar(100) NOT NULL,
    last_name varchar(100) NOT NULL,
    first_name varchar(100) NOT NULL,
    age int NOT NULL,
    CONSTRAINT vehicle_key PRIMARY KEY (id)
);

CREATE TABLE classes (
    id varchar(100) NOT NULL,
    title varchar(100) NOT NULL,
    class_description varchar(100) NOT NULL,
    CONSTRAINT person_key PRIMARY KEY (id)
);

CREATE TABLE class_students (
    student_id varchar(100) NOT NULL,
    class_id varchar(100) NOT NULL,
    CONSTRAINT student_class_key PRIMARY KEY(student_id,class_id),
    CONSTRAINT fkey_student_id FOREIGN KEY (student_id) REFERENCES students(id),
    CONSTRAINT fkey_class_id FOREIGN KEY (class_id) REFERENCES classes(id)
);

COMMIT;