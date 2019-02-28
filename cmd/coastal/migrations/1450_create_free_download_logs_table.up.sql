CREATE TABLE `free_download_logs` (
  `id`         INT UNSIGNED AUTO_INCREMENT,
  `image_id`   INT UNSIGNED NOT NUll,
  `user_id`    INT UNSIGNED NOT NUll,
  `date`       DATE,
  `created_at` DATETIME,
  PRIMARY KEY (`id`),
  UNIQUE KEY (`id`),
  INDEX date_userId (`date`, `user_id`)
)
  ENGINE = InnoDB
  CHARSET = utf8;

