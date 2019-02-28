CREATE TABLE `user_score_counts` (
  `id`                INT UNSIGNED AUTO_INCREMENT,
  `user_id`           INT UNSIGNED NOT NULL,
  `score_category_id` INT UNSIGNED NOT NULL,
  `created_at`        DATETIME,
  `updated_at`        DATETIME,
  PRIMARY KEY (`id`),
  UNIQUE KEY (`id`)
)
  ENGINE = InnoDB
  CHARSET = utf8;