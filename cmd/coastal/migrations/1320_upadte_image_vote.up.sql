ALTER TABLE `images`
  ADD `vote_up`   INT UNSIGNED DEFAULT 0 NOT NULL,
  ADD `vote_down` INT UNSIGNED DEFAULT 0 NOT NULL;