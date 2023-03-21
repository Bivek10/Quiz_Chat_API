CREATE TABLE IF NOT EXISTS chatroom (
    `id`    INT NOT NULL AUTO_INCREMENT,
    `name`  VARCHAR(200) NOT NULL,
    `created_at` DATETIME NULL,
    `updated_at` DATETIME    NULL,
    `deleted_at` DATETIME    NULL,
    PRIMARY KEY (id)
)ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;