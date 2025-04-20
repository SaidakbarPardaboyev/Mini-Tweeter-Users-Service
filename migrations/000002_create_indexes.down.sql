-- Drop indexes in reverse order

DROP INDEX IF EXISTS idx_tweet_media_tweet_id;

DROP INDEX IF EXISTS idx_tweets_user_id;

DROP INDEX IF EXISTS idx_followers_following_id;
DROP INDEX IF EXISTS idx_followers_follower_id;

DROP INDEX IF EXISTS idx_users_username;
