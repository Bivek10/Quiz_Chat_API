CREATE TABLE IF NOT EXISTS friendrequest(
    `sender`  INT NOT NULL,
    `receiver` INT NOT NULL,
    `status` VARCHAR(200) NOT NULL,
    `created_at` DATETIME NULL,
    `updated_at` DATETIME    NULL,
    `deleted_at` DATETIME    NULL,
    PRIMARY KEY (sender, receiver),
    CONSTRAINT client_sender_id_fk FOREIGN KEY (sender) REFERENCES clients (id),
    CONSTRAINT client_receiver_id_fk FOREIGN KEY (receiver) REFERENCES clients (id)
)ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;