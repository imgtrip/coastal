CREATE TABLE IF NOT EXISTS `tag`
(
  `id`         INT UNSIGNED AUTO_INCREMENT,
  `name`       VARCHAR(255) DEFAULT '' NOT NULL,
  `created_at` DATETIME,
  `updated_at` DATETIME,
  PRIMARY KEY (`id`),
  UNIQUE KEY (`id`),
  INDEX `name` (`name`)
)
  ENGINE = innodb
  CHARSET = utf8;