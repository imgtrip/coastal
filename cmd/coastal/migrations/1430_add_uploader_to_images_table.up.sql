ALTER TABLE `images`
  ADD `uploader_id` INT UNSIGNED DEFAULT 0 NOT NULL;

UPDATE `images`
SET `uploader_id` = 1799
WHERE `uploader_id` = 0;