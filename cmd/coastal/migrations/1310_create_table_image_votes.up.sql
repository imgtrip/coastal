CREATE TABLE image_votes (
  `id`         INT UNSIGNED AUTO_INCREMENT,
  `image_id`   INT UNSIGNED DEFAULT 0 NOT NULL,
  `user_id`    INT UNSIGNED DEFAULT 0 NOT NULL,
  `vote`       INT DEFAULT 0          NOT NULL,
  `created_at` DATETIME,
  `updated_at` DATETIME,
  `deleted_at` DATETIME,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  INDEX idx_imageVotes_userId_imageId(`user_id`, `image_id`)
)
  ENGINE = innoDB
  AUTO_INCREMENT = 0
  DEFAULT CHARSET = utf8;