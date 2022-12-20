package main

type Twitter struct {
	posts         []post
	subscriptions []subscriptions
}

func Constructor() Twitter {
	return Twitter{posts: make([]post, 0, 128), subscriptions: make([]subscriptions, 501)}
}

func (x *Twitter) PostTweet(userID int, tweetID int) {
	x.posts = append(x.posts, post{userID: userID, tweetID: tweetID})
}

func (x *Twitter) GetNewsFeed(userID int) []int {
	x.Follow(userID, userID)

	result := make([]int, 0, 10)
	posts, users := x.posts, x.subscriptions[userID]

	for i := len(posts) - 1; i >= 0 && len(result) < cap(result); i-- {
		if _, ok := users[posts[i].userID]; !ok {
			continue
		}

		result = append(result, posts[i].tweetID)
	}

	return result
}

func (x *Twitter) Follow(followerID int, followeeID int) {
	if x.subscriptions[followerID] == nil {
		x.subscriptions[followerID] = make(subscriptions)
	}

	x.subscriptions[followerID][followeeID] = struct{}{}
}

func (x *Twitter) Unfollow(followerID int, followeeID int) {
	if x.subscriptions[followerID] == nil {
		return
	}

	delete(x.subscriptions[followerID], followeeID)
}

type subscriptions map[int]struct{}

type post struct {
	userID  int
	tweetID int
}
