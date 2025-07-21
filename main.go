package main

import (
	"errors"
	"fmt"
	"gotorr/internal/config"
	"os"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name      string
	arguments []string
}

type commands struct {
	commands map[string]func(*state, command) error
}

func main() {
	cfg := config.Read()
	// if err := cfg.SetUser("lane"); err != nil {
	// 	fmt.Printf("setuser error: %v", err)
	// }
	// cfg = config.Read()
	s := state{
		cfg: &cfg,
	}

	commands := commands{
		commands: make(map[string]func(*state, command) error),
	}

	commands.register("login", handlerLogin)
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, "Please provide a command\n")
		os.Exit(1)
	}
	if len(os.Args) < 3 {
		fmt.Fprint(os.Stderr, "Username required\n")
		os.Exit(1)
	}

	cmd := command{
		name:      os.Args[1],
		arguments: os.Args[2:],
	}

	commands.run(&s, cmd)
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("login error: username required")
	}

	username := cmd.arguments[0]
	if err := s.cfg.SetUser(username); err != nil {
		return fmt.Errorf("login error: %w", err)
	}
	fmt.Printf("User '%s' has been set\n", username)
	return nil
}

func (c *commands) run(s *state, cmd command) error {
	if err := c.commands[cmd.name](s, cmd); err != nil {
		return err
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commands[name] = f
}
