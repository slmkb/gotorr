package main

import (
	"context"
	"fmt"
	"gotorr/internal/database"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func handlerAggregate(s *state, cmd command) error {
	// feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	// if err != nil {
	// 	return fmt.Errorf("handlerAggregate error: %w", err)
	// }
	// fmt.Println(feed)
	var time_between_reqs time.Duration
	var err error
	if len(cmd.arguments) != 0 {
		time_between_reqs, err = time.ParseDuration(cmd.arguments[0])
		if err != nil {
			log.Printf("parse time error: %v", err)
			time_between_reqs = time.Second * 30
		}
	} else {
		time_between_reqs = time.Second * 30
	}
	log.Printf("Collecting feeds every %s", time_between_reqs)
	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
	// return nil
}

func scrapeFeeds(s *state) {
	dbFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Printf("scrape feeds: %v", err)
	}
	exitId := dbFeed.ID
	var currentId uuid.UUID
	for currentId != exitId {
		dbFeed, err = s.db.GetNextFeedToFetch(context.Background())
		if err != nil {
			log.Printf("scrape feeds: %v", err)
		}
		err = s.db.MarkFeedFetched(context.Background(), dbFeed.ID)
		if err != nil {
			log.Printf("scrape feeds: %v", err)
		}
		currentId = dbFeed.ID
		urlFeed, err := fetchFeed(context.Background(), dbFeed.Url)
		if err != nil {
			log.Printf("scrape feeds: %v", err)
		}
		for _, item := range urlFeed.Channel.Item {
			pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)
			if err != nil {
				log.Printf("time parse: %v", err)
			}
			_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
				ID:          uuid.New(),
				Title:       item.Title,
				Url:         item.Link,
				Description: item.Description,
				PublishedAt: pubDate,
				FeedID:      dbFeed.ID,
			})
			if err != nil {
				log.Printf("create post: %v", err)
			}
		}
	}
}

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	var err error
	if len(cmd.arguments) != 0 {
		limit, err = strconv.Atoi(cmd.arguments[0])
		if err != nil {
			log.Printf("handler browse argument error: %v", err)
			limit = 2
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		Name:  s.cfg.CurrentUser,
		Limit: int32(limit),
	})
	if err != nil {
		log.Printf("handler browse argument error: %v", err)
		return err
	}
	for _, post := range posts {
		fmt.Println(post)
	}
	return nil
}
