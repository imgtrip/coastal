ALTER TABLE `images`
  ADD `published` BOOLEAN DEFAULT FALSE NOT NULL;

UPDATE `images`
SET `published` = TRUE;

UPDATE `images`
SET `published` = FALSE
WHERE `vote_down` >= 3;