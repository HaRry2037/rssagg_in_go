package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/harry2037/rssagg/internal/database"
)

type user struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

func databaseUserToUser(dbUser database.User) user {
	return user{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID            uuid.UUID  `json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	Name          string     `json:"name"`
	Url           string     `json:"url"`
	UserID        uuid.UUID  `json:"user_id"`
	LastFetchedAt *time.Time `json:"last_fetched_at"`
}

func databaseFeedtoFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}
func databaseFeedstoFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedtoFeed(dbFeed))
	}
	return feeds
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	UserID uuid.UUID `json:"user_id"`
	FeedID uuid.UUID `json:"feed_id"`
}

func databaseFeedFollowToFeedFollow(dbfeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbfeedFollow.FeedID,
		CreatedAt: dbfeedFollow.CreatedAt,
		UpdatedAt: dbfeedFollow.UpdatedAt,
		UserID:    dbfeedFollow.UserID,
		FeedID:    dbfeedFollow.FeedID,
	}
}

func databaseFeedFollowsToFeedFollows(dbfeedFollows []database.FeedFollow) []FeedFollow {
	feedsFollows := []FeedFollow{}
	for _, dbFeedFollow := range dbfeedFollows {
		feedsFollows = append(feedsFollows, databaseFeedFollowToFeedFollow(dbFeedFollow))
	}
	return feedsFollows
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"Title"`
	Description *string   `json:"Description"`
	PublishedAt time.Time `json:"Published_at"`
	Url         string    `json:"Url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func databasePostToPost(dbPost database.Post) Post {
	var description *string
	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}

	return Post{
		ID:          dbPost.ID,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
		Title:       dbPost.Title,
		Description: description,
		PublishedAt: dbPost.PublishedAt,
		Url:         dbPost.Url,
		FeedID:      dbPost.FeedID,
	}
}

func databasePostsToPosts(dbPost []database.Post) []Post {
	posts := []Post{}
	for _, dbPost := range dbPost {
		posts = append(posts, databasePostToPost(dbPost))

	}
	return posts
}
