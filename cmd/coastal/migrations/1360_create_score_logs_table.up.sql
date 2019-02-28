CREATE TABLE `score_logs`
(
  `id`                INT UNSIGNED AUTO_INCREMENT,
  `user_id`           INT UNSIGNED            NOT NULL,
  `score`             INT                     NOT NULL,
  `score_category_id` INT UNSIGNED            NOT NULL,
  `description`       VARCHAR(255) DEFAULT '' NOT NULL,
  `created_at`        DATETIME,

  PRIMARY KEY (`id`),
  UNIQUE KEY (`id`),
  INDEX userId_scoreCategoryId_createdAt (`user_id`, `score_category_id`,`created_at`)
)
  ENGINE = InnoDB
  CHARSET = utf8;
