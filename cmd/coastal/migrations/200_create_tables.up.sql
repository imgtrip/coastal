-- +migrate Up
SET FOREIGN_KEY_CHECKS = 0;

CREATE TABLE `image_zoom_log` (
  `id`             INT(11) NOT NULL AUTO_INCREMENT,
  `image_id`       INT(11) NOT NULL DEFAULT '0',
  `fingerprint_id` INT(11) NOT NULL DEFAULT '0',
  `created_at`     DATETIME         DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  KEY `idx_image_zoom_log_image_id` (`image_id`),
  KEY `idx_image_zoom_log_fingerprint_id` (`fingerprint_id`)
)
  ENGINE = InnoDB
  AUTO_INCREMENT = 10209
  DEFAULT CHARSET = utf8;


CREATE TABLE `album` (
  `id`              INT(11)          NOT NULL AUTO_INCREMENT,
  `title`           VARCHAR(255)     NOT NULL DEFAULT '',
  `keyword`         VARCHAR(255)     NOT NULL DEFAULT '',
  `description`     VARCHAR(255)     NOT NULL DEFAULT '',
  `brief`           TEXT,
  `user_id`         INT(11)          NOT NULL,
  `view_number`     INT(11)          NOT NULL DEFAULT '0',
  `collect_number`  INT(11)          NOT NULL DEFAULT '0',
  `created_at`      DATETIME                  DEFAULT NULL,
  `updated_at`      DATETIME                  DEFAULT NULL,
  `recommend_album` VARCHAR(255)     NOT NULL DEFAULT '',
  `visible`         TINYINT(1)                DEFAULT '0',
  `is_public`       TINYINT(1)                DEFAULT '1',
  `comments`        INT(10) UNSIGNED NOT NULL DEFAULT '0',
  `views`           INT(10) UNSIGNED NOT NULL DEFAULT '0',
  `amounts`         INT(10) UNSIGNED NOT NULL DEFAULT '0',
  `cover`           VARCHAR(255)     NOT NULL DEFAULT '',
  `score`           INT(11)          NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  AUTO_INCREMENT = 509
  DEFAULT CHARSET = utf8;


CREATE TABLE `image` (
  `id`         INT(11)      NOT NULL AUTO_INCREMENT,
  `name`       VARCHAR(255) NOT NULL DEFAULT '',
  `origin_src` TEXT,
  `origin_id`  INT(11)      NOT NULL DEFAULT '0',
  `created_at` DATETIME              DEFAULT NULL,
  `updated_at` DATETIME              DEFAULT NULL,
  `src`        VARCHAR(255) NOT NULL DEFAULT '',
  `downloads`  INT(11)      NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  AUTO_INCREMENT = 313577
  DEFAULT CHARSET = utf8;

CREATE TABLE `fingerprint` (
  `id`         INT(11)  NOT NULL AUTO_INCREMENT,
  `hash`       CHAR(32) NOT NULL DEFAULT '',
  `user_id`    INT(11)  NOT NULL DEFAULT '0',
  `created_at` DATETIME          DEFAULT NULL,
  `updated_at` DATETIME          DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  KEY `idx_fingerprint_hash` (`hash`)
)
  ENGINE = InnoDB
  AUTO_INCREMENT = 1303
  DEFAULT CHARSET = utf8;

CREATE TABLE `image_download_log` (
  `id`             INT(11) NOT NULL AUTO_INCREMENT,
  `image_id`       INT(11) NOT NULL DEFAULT '0',
  `fingerprint_id` INT(11) NOT NULL DEFAULT '0',
  `created_at`     DATETIME         DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  KEY `idx_image_download_log_image_id` (`image_id`),
  KEY `idx_image_download_log_fingerprint_id` (`fingerprint_id`)
)
  ENGINE = InnoDB
  AUTO_INCREMENT = 4875
  DEFAULT CHARSET = utf8;

CREATE TABLE `image_session` (
  `id`         INT(11)          NOT NULL AUTO_INCREMENT,
  `session_id` INT(10) UNSIGNED NOT NULL DEFAULT '0',
  `image_id`   INT(10) UNSIGNED NOT NULL DEFAULT '0',
  `created_at` DATETIME                  DEFAULT NULL,
  `updated_at` DATETIME                  DEFAULT NULL,
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  AUTO_INCREMENT = 1784568
  DEFAULT CHARSET = utf8;


CREATE TABLE `album_image` (
  `id`         INT(11) NOT NULL AUTO_INCREMENT,
  `album_id`   INT(11) NOT NULL,
  `image_id`   INT(11) NOT NULL,
  `created_at` DATETIME         DEFAULT NULL,
  `updated_at` DATETIME         DEFAULT NULL,
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  AUTO_INCREMENT = 2774
  DEFAULT CHARSET = utf8;

CREATE TABLE `comment` (
  `id`         INT(11)      NOT NULL AUTO_INCREMENT,
  `album_id`   INT(11)      NOT NULL DEFAULT '0',
  `parent_id`  INT(11)      NOT NULL DEFAULT '0',
  `user_id`    INT(11)      NOT NULL DEFAULT '0',
  `content`    VARCHAR(255) NOT NULL DEFAULT '',
  `created_at` DATETIME              DEFAULT NULL,
  `updated_at` DATETIME              DEFAULT NULL,
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  AUTO_INCREMENT = 51
  DEFAULT CHARSET = utf8;

CREATE TABLE `user` (
  `id`                  INT(11)      NOT NULL AUTO_INCREMENT,
  `name`                VARCHAR(64)  NOT NULL,
  `email`               VARCHAR(255) NOT NULL,
  `password`            VARCHAR(255) NOT NULL,
  `remember_token`      VARCHAR(255)          DEFAULT '',
  `created_at`          DATETIME              DEFAULT NULL,
  `updated_at`          DATETIME              DEFAULT NULL,
  `queue`               INT(11)      NOT NULL DEFAULT '0',
  `is_admin`            TINYINT(1)   NOT NULL DEFAULT '0',
  `notification_number` INT(11)      NOT NULL DEFAULT '0',
  `avatar`              VARCHAR(255) NOT NULL DEFAULT '',
  `album_id`            INT(11)      NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
)
  ENGINE = InnoDB
  AUTO_INCREMENT = 452
  DEFAULT CHARSET = utf8;
