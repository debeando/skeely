CREATE TABLE `Case01` (
    `id` INT(20),
    `user.name` VARCHAR(1024),
    `status` CHAR(255),
    `description` VARCHAR(2000) DEFAULT NULL,
    `update_at` datetime,
    PRIMARY KEY (`id`)
)ENGINE=MyISAM;
