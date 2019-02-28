CREATE TABLE `review_categories`
(
  `id`                         INT UNSIGNED AUTO_INCREMENT,
  `name`                       VARCHAR(64)  DEFAULT '' NOT NULL,
  `description`                VARCHAR(255) DEFAULT '' NOT NULL,
  `creator_score`              INT                     NOT NULL,
  `reviewer_score`             INT                     NOT NULL,
  `reviewer_score_required`    INT          DEFAULT 0  NOT NULL,
  `reviewer_score_category_id` INT          DEFAULT 0  NOT NULL,
  `creator_score_category_id`  INT          DEFAULT 0  NOT NULL,
  `created_at`                 DATETIME,
  `updated_at`                 DATETIME,
  PRIMARY KEY (`id`),
  UNIQUE KEY (`id`)
)
  ENGINE = InnoDB
  CHARSET = utf8;

INSERT INTO `review_categories` (`id`,
                                 `name`,
                                 `description`,
                                 `creator_score`,
                                 `reviewer_score`,
                                 `reviewer_score_required`,
                                 `reviewer_score_category_id`,
                                 `creator_score_category_id`,
                                 `created_at`,
                                 `updated_at`)
VALUES ("1", "更新图片名称", "", "5", "2", "100", "13","3","2018-12-12 00:00:00", "2018-12-12 00:00:00"),
       ("2", "创建图片标签", "", "5", "2", "100", "7","4","2018-12-12 00:00:00", "2018-12-12 00:00:00"),
       ("3", "删除图片标签", "", "2", "2", "200","8","5", "2018-12-12 00:00:00", "2018-12-12 00:00:00"),
       ("4", "删除图片", "", "2", "2", "300", "9","10","2018-12-12 00:00:00", "2018-12-12 00:00:00"),
       ("5", "删除重复图片", "", "20", "5", "400", "14","12","2018-12-12 00:00:00", "2018-12-12 00:00:00");
