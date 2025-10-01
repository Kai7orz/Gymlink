CREATE TABLE user_likes (
    `user_id` INT NOT NULL,
    `exercise_record_id` INT NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (exercise_record_id) REFERENCES exercise_records(id) ON DELETE CASCADE,
    PRIMARY KEY(user_id,exercise_record_id)
);