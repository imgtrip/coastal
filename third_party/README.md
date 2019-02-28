# `/third_party`

External helper tools, forked code and other 3rd party utilities (e.g., Swagger UI).


CREATE TABLE IF NOT EXISTS `tag`(
  `id`         int(11)     NOT NULL AUTO_INCREMENT,
  `name`       varchar(64) NOT NULL DEFAULT '',
  `created_at` datetime             DEFAULT NULL,
  `updated_at` datetime             DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  KEY `idx_tag_name` (`name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  ROW_FORMAT = FIXED;
