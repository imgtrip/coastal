-- +migrate Up
ALTER TABLE `albums`
  ADD `deleted_at` DATETIME DEFAULT NULL,
  ADD `comment_host_id` INT DEFAULT 0 NOT NULL;

ALTER TABLE `album_images`
  ADD `deleted_at` DATETIME DEFAULT NULL;

ALTER TABLE `comments`
  ADD `deleted_at` DATETIME DEFAULT NULL,
  ADD `comment_host_id` INT DEFAULT 0 NOT NULL;

ALTER TABLE `download_logs`
  ADD `deleted_at` DATETIME DEFAULT NULL,
  ADD `fingerprint` VARCHAR(64) DEFAULT '' NOT NULL;

ALTER TABLE `fingerprints`
  ADD `deleted_at` DATETIME DEFAULT NULL;

ALTER TABLE `images`
  ADD `deleted_at` DATETIME DEFAULT NULL;

ALTER TABLE `image_sessions`
  ADD `deleted_at` DATETIME DEFAULT NULL,
  ADD `session` VARCHAR(64) DEFAULT '' NOT NULL;

ALTER TABLE `users`
  ADD `deleted_at` DATETIME DEFAULT NULL;

ALTER TABLE `zoom_logs`
  ADD `deleted_at` DATETIME DEFAULT NULL,
  ADD `fingerprint` VARCHAR(64) DEFAULT '' NOT NULL;

