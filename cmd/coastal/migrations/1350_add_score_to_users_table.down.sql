ALTER TABLE `users`
  DROP COLUMN `score`,
  DROP COLUMN `referer_user_id`,
  DROP COLUMN `is_official`,
  DROP COLUMN `daily_free_download_number`;


DROP INDEX `referer_user_id`
  ON `users`;