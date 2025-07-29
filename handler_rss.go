package main

import (
	"context"
	"database/sql"
	"fmt"
	"gotorr/internal/database"
	"time"

	"github.com/google/uuid"
)

func handlerAggregate(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("handlerAggregate error: %w", err)
	}
	fmt.Println(feed)
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.arguments) < 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.name)
	}

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUser)
	if err != nil {
		return fmt.Errorf("could not retrive user: %w", err)
	}

	feedName := cmd.arguments[0]
	feedUrl := cmd.arguments[1]

	// feed, err := fetchFeed(context.Background(), feedUrl)
	// if err != nil {
	// 	return fmt.Errorf("handlerAggregate error: %w", err)
	// }

	feedParams := database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		Name: feedName,
		Url:  feedUrl,
		UserID: uuid.NullUUID{
			UUID:  currentUser.ID,
			Valid: true,
		},
	}

	feeddb, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("feed creation error: %w", err)
	}

	fmt.Println(feeddb)
	return nil
}
