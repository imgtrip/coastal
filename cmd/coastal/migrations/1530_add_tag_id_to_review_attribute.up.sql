ALTER TABLE `review_attributes`
  ADD COLUMN `tag_id`   INT UNSIGNED DEFAULT 0 NOT NULL,
  ADD COLUMN `image_id` INT UNSIGNED DEFAULT 0 NOT NULL;

ALTER TABLE `review_attributes`
  DROP COLUMN `image_tag_id`;