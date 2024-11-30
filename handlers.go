package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/soapycattt/gator/internal/database"
	"github.com/soapycattt/gator/internal/rss"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("no username is provided")
	}

	if !doesUserExist(s, cmd.args[0]) {
		return fmt.Errorf("user is not registed yet")
	}

	username := cmd.args[0]
	if err := s.cfg.SetUser(username); err != nil {
		return err
	}

	return nil
}

func handlerRemove(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("no username is provided")
	}

	if !doesUserExist(s, cmd.args[0]) {
		return fmt.Errorf("user is not registed yet")
	}

	userName := sql.NullString{
		String: cmd.args[0],
		Valid:  true,
	}
	if err := s.db.DeleteUserByName(context.Background(), userName); err != nil {
		return err
	}

	log.Printf("user %s is succesfully removed!", userName.String)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("no username is provided")
	}

	if doesUserExist(s, cmd.args[0]) {
		return fmt.Errorf("user already exists")
	}

	createdAt := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	updatedAt := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	userName := sql.NullString{
		String: cmd.args[0],
		Valid:  true,
	}

	_, err := s.db.CreateUser(
		context.Background(),
		database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			Name:      userName,
		},
	)
	if err != nil {
		return err
	}

	// if err := s.cfg.SetUser(username); err != nil {
	// 	return err
	// }

	log.Printf("user %s is succesfully registed!", userName.String)

	if err := handlerLogin(s, cmd); err != nil {
		return err
	}
	return nil
}

func handlerReset(s *state, cmd command) error {
	if err := s.db.DeleteAllUsers(context.Background()); err != nil {
		return err
	}
	log.Println("all users have been removed sucessfully!")

	if err := s.db.DeleteAllFeeds(context.Background()); err != nil {
		return err
	}
	log.Println("all feeds have been removed sucessfully!")

	return nil
}

func handlerList(s *state, cmd command) error {
	users, err := s.db.ListAllUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		output := "* " + user.Name.String
		if user.Name.String == s.cfg.CurrentUser {
			output += " (current)"
		}

		fmt.Println(output)
	}

	return nil
}

func handlerAgg(s *state, cmd command) error {
	url := "https://www.wagslane.dev/index.xml"
	rssFeed, err := rss.FetchFeed(context.Background(), url)
	if err != nil {
		return err
	}

	for _, item := range rssFeed.Channel.Item {
		fmt.Println(" - " + item.Title)
	}

	return nil
}

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) <= 1 {
		return fmt.Errorf("expect 2 arugments, got %v", len(cmd.args))
	}

	feedName := cmd.args[0]
	feedURL := cmd.args[1]

	rssFeed, err := rss.FetchFeed(context.Background(), feedURL)
	if err != nil {
		return err
	}

	fmt.Printf("User: %s\n", user.Name.String)
	fmt.Printf("Feed: %s\n", feedName)

	feed, err := s.db.CreateFeeds(
		context.Background(),
		database.CreateFeedsParams{
			ID:        uuid.New(),
			CreatedAt: sql.NullTime{time.Now(), true},
			UpdatedAt: sql.NullTime{time.Now(), true},
			Name:      sql.NullString{rssFeed.Channel.Title, true},
			Url:       sql.NullString{rssFeed.Channel.Link, true},
			UserID:    uuid.NullUUID{user.ID, true},
			RssUrl:    sql.NullString{feedURL, true},
		},
	)
	if err != nil {
		return err
	}

	_, err = s.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: sql.NullTime{time.Now(), true},
			UpdatedAt: sql.NullTime{time.Now(), true},
			UserID:    uuid.NullUUID{user.ID, true},
			FeedID:    uuid.NullUUID{feed.ID, true},
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.ListFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Printf(
			" - Name: %s\n\t-Owner: %s\n\t-URL: %s\n",
			feed.Feed.String,
			feed.User.String,
			feed.Url.String,
		)
	}

	return nil
}

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("no url is provided")
	}

	feedUrl := cmd.args[0]

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

	feed, err := s.db.GetFeedByURL(
		context.Background(),
		sql.NullString{
			feedUrl,
			true,
		},
	)
	if err != nil {
		return err
	}

	_, err = s.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: sql.NullTime{time.Now(), true},
			UpdatedAt: sql.NullTime{time.Now(), true},
			UserID:    uuid.NullUUID{user.ID, true},
			FeedID:    uuid.NullUUID{feed.ID, true},
		},
	)
	if err != nil {
		return err
	}

	fmt.Printf("%s followed %s", user.Name.String, feed.Name.String)

	return nil
}

func handlerFollowing(s *state, cmd command) error {
	followings, err := s.db.GetFeedFollowsForUser(
		context.Background(),
		sql.NullString{s.cfg.CurrentUser, true},
	)
	if err != nil {
		return err
	}

	for _, following := range followings {
		fmt.Printf("- %s\n", following.String)
	}

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("no url is provided")
	}
	feedUrl := cmd.args[0]

	feed, err := s.db.GetFeedByURL(
		context.Background(),
		sql.NullString{
			feedUrl,
			true,
		},
	)
	if err != nil {
		return err
	}

	err = s.db.DeleteFollow(
		context.Background(),
		database.DeleteFollowParams{
			uuid.NullUUID{user.ID, true},
			uuid.NullUUID{feed.ID, true},
		},
	)
	if err != nil {
		return err
	}

	fmt.Printf("%s has just unfollowed %s", user.Name.String, feed.Name.String)

	return nil
}

func doesUserExist(s *state, userName string) bool {
	user, _ := s.db.GetUserByName(context.Background(), sql.NullString{
		String: userName,
		Valid:  true,
	})
	return user.ID != uuid.Nil
}
