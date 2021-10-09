package twitter

// TweetRecentCountsResponse contains all of the information from a tweet recent counts
type TweetRecentCountsResponse struct {
	Raw  *TweetCountsRaw
	Meta *TweetRecentCountsMeta `json:"meta"`
}

// TweetRecentCountsMeta contains the meta data from the recent counts information
type TweetRecentCountsMeta struct {
	TotalTweetCount int `json:"total_tweet_count"`
}

// TweetCountsRaw is the raw response from the tweet counts endpoint
type TweetCountsRaw struct {
	TweetCounts []*TweetCount `json:"data"`
}

// TweetCount is the object on the tweet counts endpoints
type TweetCount struct {
	Start      string `json:"start"`
	End        string `json:"end"`
	TweetCount int    `json:"tweet_count"`
}
