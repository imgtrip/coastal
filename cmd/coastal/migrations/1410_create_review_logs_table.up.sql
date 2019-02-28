CREATE TABLE `review_logs` (
  `id`         INT UNSIGNED AUTO_INCREMENT,
  `review_id`  INT UNSIGNED NOT NULL,
  `user_id`    INT UNSIGNED NOT NULL,
  `passed`     BOOLEAN      NOT NULL,
  `created_at` DATETIME,
  PRIMARY KEY (`id`),
  UNIQUE KEY (`id`)
)
  ENGINE = InnoDB
  CHARSET = utf8;
