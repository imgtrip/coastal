-- +migrate Up
RENAME TABLE
    `album` TO `albums`,
    `album_image` TO `album_images`,
    `comment` TO `comments`,
    `image_download_log` TO `download_logs`,
    `fingerprint` TO `fingerprints`,
    `image` TO `images`,
    `image_session` TO `image_sessions`,
    `user` TO `users`,
    `image_zoom_log` TO `zoom_logs`;
