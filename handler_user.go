package main

import (
	"context"
	"database/sql"
	"fmt"
	"gotorr/internal/database"
	"log"
	"time"

	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("usage: %s <user>", cmd.name)
	}

	username := cmd.arguments[0]
	if err := s.cfg.SetUser(username); err != nil {
		return fmt.Errorf("could not set current user: %w", err)
	}

	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("could not retrive user: %w", err)
	}

	fmt.Printf("User %s logged in successfully\n", username)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("usage: %s <user>", cmd.name)
	}

	username := cmd.arguments[0]
	createUserParams := database.CreateUserParams{
		ID: uuid.New(),

		CreatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		Name: username,
	}

	dbu, err := s.db.CreateUser(context.Background(), createUserParams)
	if err != nil {
		return fmt.Errorf("user creation error: %w", err)
	}

	if err := s.cfg.SetUser(username); err != nil {
		return fmt.Errorf("set user error: %w", err)
	}

	fmt.Printf("User %s successfully created\n", dbu.Name)
	log.Printf("database user created: %+v", dbu)
	return nil
}

func handlerReset(s *state, cmd command) error {
	if err := s.db.DeleteUsers(context.Background()); err != nil {
		return fmt.Errorf("reset handler error: %w", err)
	}
	fmt.Println("Database reset successfull")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("get users error: %w", err)
	}

	currentUser := s.cfg.CurrentUser
	for _, name := range users {
		if name == currentUser {
			fmt.Printf("* %s (current)\n", name)
		} else {
			fmt.Printf("* %s\n", name)
		}
	}
	return nil
}
