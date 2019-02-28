ALTER TABLE `albums`
  DROP COLUMN `keyword`,
  DROP COLUMN `description`,
  DROP COLUMN `brief`,
  DROP COLUMN `view_number`,
  DROP COLUMN `collect_number`,
  DROP COLUMN `recommend_album`,
  DROP COLUMN `visible`;

ALTER TABLE `comments`
  DROP COLUMN `parent_id`;

ALTER TABLE `users`
  DROP COLUMN `remember_token`,
  DROP COLUMN `queue`,
  DROP COLUMN `is_admin`,
  DROP COLUMN `notification_number`;

alter table `image_sessions`
  DROP COLUMN `session_id`;

ALTER TABLE `zoom_logs`
  DROP COLUMN `fingerprint_id`;