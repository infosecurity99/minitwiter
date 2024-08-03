
alter table users
    drop column if exists created_at,
    drop column if exists updated_at,
    drop column if exists deleted_at;

alter table tweets
    drop column if exists created_at,
    drop column if exists updated_at,
    drop column if exists deleted_at;

alter table followers
    drop column if exists created_at,
    drop column if exists updated_at,
    drop column if exists deleted_at;

alter table likes
    drop column if exists created_at,
    drop column if exists updated_at,
    drop column if exists deleted_at;

alter table retweets
    drop column if exists created_at,
    drop column if exists updated_at,
    drop column if exists deleted_at;

