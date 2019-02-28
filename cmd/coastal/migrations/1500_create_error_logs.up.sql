create table `error_logs`
(
  `id`          int unsigned auto_increment,
  `code`        int unsigned default 0  not null,
  `message`     varchar(255) default '' not null,
  `url`         varchar(255) default '' not null,
  `payload`     text,
  `environment` varchar(255) default '' not null,
  `header`      text,
  `cookie`      text,
  `created_at`  datetime,
  `updated_at`  datetime,
  primary key (`id`),
  unique key (`id`)
)
  ENGINE = InnoDB
  CHARSET = utf8;
