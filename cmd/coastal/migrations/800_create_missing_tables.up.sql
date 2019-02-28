-- +migrate Up
DROP TABLE IF EXISTS `comment_hosts`;
CREATE TABLE `comment_hosts` (
  `id`         INT(11) NOT NULL AUTO_INCREMENT,
  `created_at` DATETIME         DEFAULT NULL,
  `updated_at` DATETIME         DEFAULT NULL,
  `deleted_at` DATETIME         DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
)
  ENGINE = InnoDB
  AUTO_INCREMENT = 0
  DEFAULT CHARSET = utf8;


DROP TABLE IF EXISTS `posts`;
CREATE TABLE `posts` (
  `id`              INT(11)                 NOT NULL AUTO_INCREMENT,
  `title`           VARCHAR(255) DEFAULT '' NOT NULL,
  `user_id`         INT(11) DEFAULT 0       NOT NULL,
  `comments`        INT(11) DEFAULT 0       NOT NULL,
  `views`           INT(11) DEFAULT 0       NOT NULL,
  `comment_host_id` INT(11) DEFAULT 0       NOT NULL,
  `body`            TEXT,
  `cover`           VARCHAR(255) DEFAULT '' NOT NULL,
  `created_at`      DATETIME                         DEFAULT NULL,
  `updated_at`      DATETIME                         DEFAULT NULL,
  `deleted_at`      DATETIME                         DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
)
  ENGINE = InnoDB
  AUTO_INCREMENT = 0
  DEFAULT CHARSET = utf8;


DROP TABLE IF EXISTS `tokens`;
CREATE TABLE `tokens` (
  `id`         INT(11)                      NOT NULL AUTO_INCREMENT,
  `title`      VARCHAR(255) DEFAULT ''      NOT NULL,
  `user_id`    INT(11) DEFAULT 0            NOT NULL,
  `hash`       VARCHAR(32) DEFAULT ''       NOT NULL,
  `expire_at`  DATETIME                              DEFAULT NULL,
  `created_at` DATETIME                              DEFAULT NULL,
  `updated_at` DATETIME                              DEFAULT NULL,
  `deleted_at` DATETIME                              DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
)
  ENGINE = InnoDB
  AUTO_INCREMENT = 0
  DEFAULT CHARSET = utf8;