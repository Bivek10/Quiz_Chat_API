CREATE TABLE
    IF NOT EXISTS chatmessage (
        `id` INT NOT NULL AUTO_INCREMENT,
        `message` TEXT NOT NULL,
        `room_id` INT NOT NULL,
        `user_id` INT NOT NULL,
        `created_at` DATETIME NULL,
        `updated_at` DATETIME NULL,
        `deleted_at` DATETIME NULL,
        PRIMARY KEY (id),
        CONSTRAINT message_room_id_fk FOREIGN KEY (room_id) REFERENCES chatroom (id),
        CONSTRAINT message_user_id_fk FOREIGN KEY (user_id) REFERENCES clients (id)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
    

