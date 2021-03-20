package main

type Twitter struct {
	tweets       []int
	tweetsAuthor map[int]int
	following    map[int]map[int]bool
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{following: make(map[int]map[int]bool), tweetsAuthor: make(map[int]int)}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.tweets = append(this.tweets, tweetId)
	this.tweetsAuthor[tweetId] = userId
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	var result []int
	for i := len(this.tweets) - 1; i >= 0 && len(result) < 10; i-- {
		tweetId := this.tweets[i]
		tweetAuthor := this.tweetsAuthor[tweetId]
		if tweetAuthor == userId || (this.following[userId] != nil && this.following[userId][tweetAuthor]) {
			result = append(result, tweetId)
		}
	}
	return result
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	if this.following[followerId] == nil {
		this.following[followerId] = make(map[int]bool)
	}
	this.following[followerId][followeeId] = true
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	if this.following[followerId] == nil {
		return
	}
	delete(this.following[followerId], followeeId)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
