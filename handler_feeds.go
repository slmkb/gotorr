package main

import (
	"context"
	"fmt"
	"gotorr/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("getfeeds: %w", err)
	}

	for _, feed := range feeds {
		username, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("getuserbyid: %w", err)
		}
		fmt.Printf("Name: %s, URL: %s, Username: %s\n", feed.Name, feed.Url, username)
	}
	return nil
}

func handlerFollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("usage: %s <url>", cmd.name)
	}

	url := cmd.arguments[0]

	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("could not retrive feed: %w", err)
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	feedFollowRow, err := s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("could not create feed follow: %w", err)
	}

	fmt.Printf("Feed: %s\n", feedFollowRow.FeedName)
	fmt.Printf("User: %s\n", feedFollowRow.UserName)
	return nil
}

func handlerGetFeeds(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return fmt.Errorf("could not get feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Println(feed)
	}
	return nil
}

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.name)
	}

	feedName := cmd.arguments[0]
	feedUrl := cmd.arguments[1]

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedUrl,
		UserID:    user.ID,
	}

	feeddb, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("feed creation error: %w", err)
	}

	feedFolowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: feeddb.CreatedAt,
		UpdatedAt: feeddb.UpdatedAt,
		UserID:    feeddb.UserID,
		FeedID:    feedParams.ID,
	}
	_, err = s.db.CreateFeedFollow(context.Background(), feedFolowParams)
	if err != nil {
		return fmt.Errorf("feed follow creation error: %w", err)
	}
	return nil
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {

	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUser)
		if err != nil {
			return fmt.Errorf("mwli: %v", err)
		}
		return handler(s, cmd, user)
	}
}

func handlerUnfollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 1 {
		return fmt.Errorf("usage: %s <url>", cmd.name)
	}

	dat := database.UnfollowFeedParams{
		UserID: user.ID,
		Url:    cmd.arguments[0],
	}
	err := s.db.UnfollowFeed(context.Background(), dat)
	if err != nil {
		return fmt.Errorf("unfollow feed: %w", err)
	}
	return nil
}
