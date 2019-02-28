DROP PROCEDURE IF EXISTS sync_album_comment_host_id;
CREATE PROCEDURE sync_album_comment_host_id()
  BEGIN
    DECLARE v_counter INT UNSIGNED DEFAULT 0;
    SELECT COUNT(*)
        INTO @albumRows FROM albums;

    START TRANSACTION;
    WHILE v_counter < @albumRows DO
      INSERT INTO comment_hosts (created_at, updated_at) VALUES (NOW(), NOW());
      SELECT LAST_INSERT_ID()
          INTO @lastCommentHostId;

      UPDATE albums SET comment_host_id = @lastCommentHostId WHERE comment_host_id = 0 LIMIT 1;

      update comments
      set comment_host_id = @lastCommentHostId
      where album_id = (select id from albums where comment_host_id = @lastCommentHostId);

      SET v_counter = v_counter + 1;
    END WHILE;
    COMMIT;
  END;
CALL sync_album_comment_host_id();

