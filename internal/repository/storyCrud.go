package repository

import (
	"fmt"
	"strings"

	"github.com/isido5ik/StoryPublishingPlatform/dtos"
	"github.com/sirupsen/logrus"
)

func (r *repository) CreateStory(story dtos.AddPostInput, userId int) (int, error) {
	var postId int

	query := fmt.Sprintf("INSERT INTO %s (user_id, title, content ) VALUES($1, $2, $3) RETURNING post_id", postsTable)
	row := r.db.QueryRow(query, userId, story.Title, story.Content)
	if err := row.Scan(&postId); err != nil {
		logrus.WithError(err).Error("Failed to Scan postId")
		return 0, err
	}
	return postId, nil
}

func (r *repository) GetStories(pagination dtos.PaginationParams) ([]dtos.Post, error) {
	var stories []dtos.Post

	offset := (pagination.Page - 1) * pagination.PageSize

	query := fmt.Sprintf(
		"SELECT u.username, p.user_id, p.post_id, p.title, p.content, p.comments, p.likes, p.created_at FROM %s p JOIN %s u ON p.user_id = u.user_id LIMIT %d OFFSET %d",
		postsTable, usersTable, pagination.PageSize, offset)

	if err := r.db.Select(&stories, query); err != nil {
		logrus.WithError(err).Error("Failed to select all stories")
		return nil, err
	}

	return stories, nil
}

func (r *repository) GetUsersStories(userId int) (string, []dtos.Post, error) {
	var username string
	var stories []dtos.Post

	getStoriesQuery := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", postsTable)

	if err := r.db.Select(&stories, getStoriesQuery, userId); err != nil {
		logrus.WithError(err).Error("Failed to select users stories")
		return "", nil, err
	}
	return username, stories, nil
}

func (r *repository) GetStory(postId int) (dtos.Post, error) {
	var story dtos.Post

	query := fmt.Sprintf("SELECT u.username, p.user_id, p.post_id, p.title, p.content, p.comments, p.likes, p.created_at  FROM %s p JOIN %s u ON p.user_id = u.user_id WHERE post_id = $1", postsTable, usersTable)
	if err := r.db.Get(&story, query, postId); err != nil {
		logrus.WithError(err).Error("Failed to get story by id")
		return story, err
	}

	return story, nil
}

func (r *repository) DeleteStory(postId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE post_id = $1", postsTable)
	_, err := r.db.Exec(query, postId)
	if err != nil {
		logrus.WithError(err).Error("Failed to delete story")
		return err
	}

	return nil
}

func (r *repository) UpdateStory(postId int, input dtos.UpdateStoryInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Content != "" {
		setValues = append(setValues, fmt.Sprintf("content=$%d", argId))
		args = append(args, input.Content)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE post_id = %d", postsTable, setQuery, postId)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		logrus.WithError(err).Error("Failed to update story")
		return err
	}
	return nil
}
