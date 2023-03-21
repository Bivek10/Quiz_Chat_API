
CREATE TABLE IF NOT EXISTS conversations (
    `id`    INT NOT NULL AUTO_INCREMENT,
    `sender`  INT NOT NULL,
    `receiver` INT NOT NULL,
    `message` TEXT NOT NULL,
    `status`  VARCHAR(50) NOT NULL,
    `created_at` DATETIME NULL,
    `updated_at` DATETIME    NULL,
    `deleted_at` DATETIME    NULL,
    PRIMARY KEY (id),
    CONSTRAINT client_senders_id_fk FOREIGN KEY (sender) REFERENCES clients (id),
    CONSTRAINT client_receivers_id_fk FOREIGN KEY (receiver) REFERENCES clients (id)
)ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;