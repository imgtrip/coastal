ALTER TABLE `review_attributes`
  ADD COLUMN `image_tag_id` INT UNSIGNED DEFAULT 0 NOT NULL;

ALTER TABLE `review_attributes`
  DROP COLUMN `tag_id`,
  DROP COLUMN `image_id`;