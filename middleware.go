package main

import (
	"context"
	"database/sql"

	"github.com/soapycattt/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUserByName(
			context.Background(),
			sql.NullString{
				s.cfg.CurrentUser,
				true,
			},
		)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
