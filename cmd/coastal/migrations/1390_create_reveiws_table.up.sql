CREATE TABLE `reviews` (
  `id`                 INT UNSIGNED AUTO_INCREMENT,
  `creator_id`         INT UNSIGNED           NOT NULL,
  `image_id`           INT UNSIGNED           NOT NULL,
  `review_category_id` INT UNSIGNED           NOT NULL,
  `agree_count`        INT UNSIGNED DEFAULT 0 NOT NULL,
  `disagree_count`     INT UNSIGNED DEFAULT 0 NOT NULL,
  `end_at`             DATETIME,
  `created_at`         DATETIME,
  `updated_at`         DATETIME,
  `deleted_at`         DATETIME,
  PRIMARY KEY (`id`),
  UNIQUE KEY (`id`),
  INDEX `creator_image_category` (`creator_id`, `image_id`, `review_category_id`)
)
  ENGINE = InnoDB
  CHARSET = utf8;