package main

import (
	"fmt"

	"github.com/soapycattt/gator/internal/config"
	"github.com/soapycattt/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	mapping map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.mapping[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	cmdFunc, found := c.mapping[cmd.name]
	if !found {
		return fmt.Errorf("no function %v found \n args are %v", cmd.name, cmd.args)
	}

	if err := cmdFunc(s, cmd); err != nil {
		return err
	}
	return nil
}
