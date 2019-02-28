CREATE TABLE verification_codes (
  `id`         INT(11)     NOT NULL AUTO_INCREMENT,
  `email`      VARCHAR(64) NOT NULL DEFAULT '',
  `code`       VARCHAR(10) NOT NULL DEFAULT '',
  `token`      VARCHAR(32) NOT NULL DEFAULT '',
  `created_at` DATETIME             DEFAULT NULL,
  `updated_at` DATETIME             DEFAULT NULL,
  `deleted_at` DATETIME             DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  KEY `idx_verificationCodes_email_code` (`email`, `code`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8;
