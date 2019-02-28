create table `image_tag_vote_logs`
(
  `id`           int unsigned auto_increment,
  `image_id`     INT UNSIGNED DEFAULT 0 NOT NULL,
  `image_tag_id` int unsigned default 0 not null,
  `user_id`      int unsigned default 0 not null,
  `vote`         int          default 0 not null,
  `created_at`   datetime,
  primary key (`id`),
  unique key (`id`),
  index imageTagId_userId (`image_tag_id`, `user_id`)
)
  ENGINE = InnoDB
  CHARSET = utf8;