CREATE TABLE `review_attributes` (
  `id`             INT UNSIGNED AUTO_INCREMENT,
  `review_id`      INT UNSIGNED            NOT NULL,
  `image_tag_name` VARCHAR(255) DEFAULT "" NOT NUll,
  `image_name`     VARCHAR(255) DEFAULT "" NOT NUll,
  `image_tag_id`   INT UNSIGNED DEFAULT 0  NOT NUll,
  `created_at`     DATETIME,
  `updated_at`     DATETIME,
  PRIMARY KEY (`id`),
  UNIQUE KEY (`id`),
  INDEX reviewId (`review_id`)
)
  ENGINE = InnoDB
  CHARSET = utf8;

