package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/soapycattt/gator/internal/config"
	"github.com/soapycattt/gator/internal/database"
)

func main() {
	/*
		- In the main function, remove the manual update of the config file. Instead, simply read the config file, and store the config in a new instance of the state struct.
		- Create a new instance of the commands struct with an initialized map of handler functions.
		- Register a handler function for the login command.
		- Use os.Args to get the command-line arguments passed in by the user.
	*/

	// Read config
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	dbQueries := database.New(db)

	// Init vars
	s := state{dbQueries, cfg}
	c := commands{
		make(map[string]func(*state, command) error),
	}

	// Register handler
	c.register("login", handlerLogin)
	c.register("register", handlerRegister)
	c.register("remove", handlerRemove)
	c.register("reset", handlerReset)
	c.register("users", handlerList)
	c.register("list", handlerList)
	c.register("agg", handlerAgg)
	c.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	c.register("feeds", handlerFeeds)
	c.register("follow", handlerFollow)
	c.register("following", handlerFollowing)
	c.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	c.register("browse", middlewareLoggedIn(handlerBrowse))

	// Get inputArgs
	inputArgs := os.Args

	if len(inputArgs) < 2 {
		log.Fatalf("expect at least 1 argument, got 0")
		os.Exit(1)
	}

	// Run the command
	cmdName := inputArgs[1]
	args := inputArgs[2:]
	cmd := command{cmdName, args}

	if err := c.run(&s, cmd); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	os.Exit(0)
}
