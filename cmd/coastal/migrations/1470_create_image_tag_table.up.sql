CREATE TABLE IF NOT EXISTS `image_tag`
(
  `id`         INT UNSIGNED AUTO_INCREMENT,
  `image_id`   INT UNSIGNED NOT NUll,
  `tag_id`     INT UNSIGNED NOT NUll,
  `confidence` INT UNSIGNED DEFAULT 0 NOT NULL,
  `created_at` DATETIME,
  `updated_at` DATETIME,
  PRIMARY KEY (`id`),
  UNIQUE KEY (`id`),
  INDEX imageId(`image_id`)
)
  ENGINE = InnoDB
  CHARSET = utf8;