CREATE TABLE comment (
    id VARCHAR(20) PRIMARY KEY,
    context text NOT NULL,
    video_id  VARCHAR(20) NOT NULL,
    channel_id VARCHAR(20) NOT NULL
);

-- name: CreateComment :one
INSERT INTO comment (
  id, context, video_id, channel_id
) VALUES (
  $1, $2, $3, $4
);