package main

import (
	"database/sql"
	"gotorr/internal/config"
	"gotorr/internal/database"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config file: %v\n", err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatalf("Database initialization error: %v\n", err)
	}
	defer db.Close()

	dbQueries := database.New(db)

	s := state{
		cfg: &cfg,
		db:  dbQueries,
	}

	commands := commands{
		regsiteredCommands: make(map[string]func(*state, command) error),
	}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)
	commands.register("agg", handlerAggregate)
	commands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	commands.register("feeds", handlerListFeeds)
	commands.register("follow", middlewareLoggedIn(handlerFollowFeed))
	commands.register("following", middlewareLoggedIn(handlerGetFeeds))
	commands.register("unfollow", middlewareLoggedIn(handlerUnfollowFeed))
	commands.register("browse", middlewareLoggedIn(handlerBrowse))

	if len(os.Args) < 2 {
		log.Fatalf("Please provide a command\n")
	}

	cmd := command{
		name:      os.Args[1],
		arguments: os.Args[2:],
	}

	if err := commands.run(&s, cmd); err != nil {
		log.Fatalf("Command error: %v\n", err)
	}
}
