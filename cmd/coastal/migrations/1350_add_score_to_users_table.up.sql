ALTER TABLE `users`
  ADD `score`                      INT UNSIGNED DEFAULT 0 NOT NULL,
  ADD `referer_user_id`            INT UNSIGNED DEFAULT 0 NOT NULL,
  ADD `is_official`                BOOL DEFAULT FALSE NOT NULL,
  ADD `daily_free_download_number` INT UNSIGNED DEFAULT 0 NOT NULL;

CREATE INDEX `referer_user_id`
  ON `users` (`referer_user_id`);