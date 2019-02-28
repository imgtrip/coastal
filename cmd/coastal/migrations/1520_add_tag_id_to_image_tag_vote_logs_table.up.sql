alter table `image_tag_vote_logs`
  add `tag_id` int unsigned default 0 not null;


DROP INDEX `imageTagId_userId` ON `image_tag_vote_logs`;

create index `imageId_tagId_userId` on `image_tag_vote_logs` (`image_id`, `tag_id`, `user_id`);
create index `imageTagId` on `image_tag_vote_logs` (`image_tag_id`);
