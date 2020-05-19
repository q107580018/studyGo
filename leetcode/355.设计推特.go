package main

import "fmt"

func main() {
	t := Constructor()
	t.PostTweet(1, 5)
	t.Follow(1, 1)
	fmt.Println(t.GetNewsFeed(1))
}

type Twitter struct {
	tweets     [][2]int
	follows    map[int][]int
	isFollowMe map[int]bool
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	var tweets = make([][2]int, 0, 10)
	var follows = make(map[int][]int)
	var isFollowMe = make(map[int]bool)
	return Twitter{
		tweets:     tweets,
		follows:    follows,
		isFollowMe: isFollowMe,
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.tweets = append(this.tweets, [2]int{userId, tweetId})
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	News := make([]int, 0, 10)
	follows := this.follows[userId]
	if !this.isFollowMe[userId] {
		follows = append(follows, userId)
	}
	count := 0
	for i := len(this.tweets) - 1; i >= 0 && count < 10; i-- {
		for _, id := range follows {
			if id == this.tweets[i][0] {
				News = append(News, this.tweets[i][1])
				count++
				break
			}
		}
	}
	return News
}

///** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	for _, id := range this.follows[followerId] {
		if id == followeeId {
			return
		}
	}
	if followerId == followeeId {
		this.isFollowMe[followerId] = true
	}
	this.follows[followerId] = append(this.follows[followerId], followeeId)
}

///** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	for i, id := range this.follows[followerId] {
		if id == followeeId {
			this.follows[followerId] = append(this.follows[followerId][:i], this.follows[followerId][i+1:]...)
			if followeeId == followerId {
				this.isFollowMe[followerId] = false
			}
			return
		}

	}

}
