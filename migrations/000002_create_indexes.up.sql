-- Users indexes
CREATE INDEX idx_users_username ON users(username);

-- Followers indexes
CREATE INDEX idx_followers_follower_id ON followers(follower_id);
CREATE INDEX idx_followers_following_id ON followers(following_id);

-- Tweets indexes
CREATE INDEX idx_tweets_user_id ON tweets(user_id);

-- Tweet media index
CREATE INDEX idx_tweet_media_tweet_id ON tweet_media(tweet_id);