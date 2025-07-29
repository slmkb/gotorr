package main

import (
	"context"
	"fmt"
)

func handlerListFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("getfeeds: %w", err)
	}

	for _, feed := range feeds {
		username, err := s.db.GetUserByID(context.Background(), feed.UserID.UUID)
		if err != nil {
			return fmt.Errorf("getuserbyid: %w", err)
		}
		fmt.Printf("Name: %s, URL: %s, Username: %s\n", feed.Name, feed.Url, username)
	}
	return nil
}
