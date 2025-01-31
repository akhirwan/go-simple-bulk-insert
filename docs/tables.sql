CREATE TABLE `test_number_counter` (
`type` CHAR(4) NOT NULL COLLATE 'utf8mb4_unicode_ci',
`start_number` INT(10) NOT NULL,
`last_number` INT(10) NULL DEFAULT NULL,
PRIMARY KEY (`type`) USING BTREE
)
COLLATE='utf8mb4_unicode_ci'
ENGINE=InnoDB;

INSERT INTO `test_number_counter` (`type`, `start_number`, `last_number`)
VALUES ('TP1', 8000000, NULL);
INSERT INTO `test_number_counter` (`type`, `start_number`, `last_number`)
VALUES ('TP2', 1000000, NULL);

CREATE TABLE `test_number_transaction` (
`number` INT(10) NOT NULL,
`action` VARCHAR(50) NOT NULL COLLATE 'utf8mb4_unicode_ci',
PRIMARY KEY (`number`) USING BTREE,
UNIQUE INDEX `number` (`number`) USING BTREE
)
COLLATE='utf8mb4_unicode_ci'
ENGINE=InnoDB;