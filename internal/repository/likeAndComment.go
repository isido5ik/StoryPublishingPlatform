package repository

import (
	"fmt"
	"log"

	"github.com/isido5ik/StoryPublishingPlatform/dtos"
	"github.com/sirupsen/logrus"
)

func (r *repository) CheckLike(userId, postId int) error {

	var like dtos.Like
	checkLikeQuery := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1 AND post_id = $2", likesTable)
	err := r.db.Get(&like, checkLikeQuery, userId, postId)
	logrus.Info(checkLikeQuery, userId, postId)
	return err
}

func (r *repository) CheckComment(userId, postId, commentId int) error {
	var comment dtos.Comment
	checkCommentQuery := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1 AND post_id = $2 AND comment_id = $3", commentsTable)
	err := r.db.Get(&comment, checkCommentQuery, userId, postId, commentId)
	return err
}

func (r *repository) Like(userId, postId int) (int, error) {

	err := r.CheckLike(userId, postId)
	if err == nil {
		//already liked
		return -1, err
	}

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	incLikeQuery := fmt.Sprintf("UPDATE %s SET likes = likes + 1 WHERE post_id = $1", postsTable)
	_, err = tx.Exec(incLikeQuery, postId)
	if err != nil {
		logrus.WithError(err).Error("Failed to increment likes")
		tx.Rollback()
		return 0, err
	}

	var likeId int
	addLikeQuery := fmt.Sprintf("INSERT INTO %s (user_id, post_id) VALUES($1, $2) RETURNING like_id", likesTable)
	row := r.db.QueryRow(addLikeQuery, userId, postId)
	if err := row.Scan(&likeId); err != nil {
		logrus.WithError(err).Error("Failed to add like")
		tx.Rollback()
		return 0, err
	}

	return likeId, tx.Commit()

}

func (r *repository) RemoveLike(userId, postId int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	decLikeQuery := fmt.Sprintf("UPDATE %s SET likes = likes - 1 WHERE post_id = $1", postsTable)
	_, err = tx.Exec(decLikeQuery, postId)
	if err != nil {
		logrus.WithError(err).Error("Failed to decrement likes")
		tx.Rollback()
		return err
	}
	removeLikeQuery := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1", likesTable)
	_, err = tx.Exec(removeLikeQuery, userId)
	if err != nil {
		logrus.WithError(err).Error("Failed to remove like")
		tx.Rollback()
		return err
	}

	return tx.Commit()

}

func (r *repository) AddComment(userId, postId int, comment dtos.AddCommentInput) (int, error) {
	log.Print(userId, postId, comment.CommentText)
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	incCommentQuery := fmt.Sprintf("UPDATE %s SET comments = comments + 1 WHERE post_id = $1", postsTable)
	_, err = tx.Exec(incCommentQuery, postId)
	if err != nil {
		logrus.WithError(err).Error("Failed to increment comments")
		tx.Rollback()
		return 0, err
	}

	var commentId int

	addCommentQuery := fmt.Sprintf(`
		INSERT INTO %s (comment_text, user_id, post_id) 
		VALUES ($1, $2, $3) 
		RETURNING comment_id`, commentsTable)

	row := r.db.QueryRow(addCommentQuery, comment.CommentText, userId, postId)
	if err := row.Scan(&commentId); err != nil {
		logrus.WithError(err).Error("Failed to add comment")
		tx.Rollback()
		return 0, err
	}
	return commentId, tx.Commit()
}

func (r *repository) ReplyToComment(userId, postId, parentId int, comment dtos.AddCommentInput) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	incCommentQuery := fmt.Sprintf("UPDATE %s SET comments = comments + 1 WHERE post_id = $1", postsTable)
	_, err = tx.Exec(incCommentQuery, postId)
	if err != nil {
		logrus.WithError(err).Error("Failed to increment comments")
		tx.Rollback()
		return 0, err
	}
	var commentId int

	addCommentQuery := fmt.Sprintf(`
		INSERT INTO %s (comment_text, user_id, post_id, parent_id) 
		VALUES ($1, $2, $3, $4) 
		RETURNING comment_id`, commentsTable)

	row := r.db.QueryRow(addCommentQuery, comment.CommentText, userId, postId, parentId)
	if err := row.Scan(&commentId); err != nil {
		logrus.WithError(err).Error("Failed to add comment (reply)")
		tx.Rollback()
		return 0, err
	}
	return commentId, tx.Commit()

}

func (r *repository) UpdateComment(userId, postId, commentId int, newComment dtos.UpdateCommentInput) error {
	updateCommentQuery := fmt.Sprintf("UPDATE %s SET comment_text = $1 WHERE user_id = $2 AND post_id = $3 AND comment_id = $4", commentsTable)
	_, err := r.db.Exec(updateCommentQuery, newComment.CommentText, userId, postId, commentId)
	if err != nil {
		logrus.WithError(err).Error("Failed to update comment")
		return err
	}
	return nil
}

func (r *repository) DeleteComment(userId, postId, commentId int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	decCommentQuery := fmt.Sprintf("UPDATE %s SET comments = comments - 1 WHERE post_id = $1 AND user_id = $2", postsTable)
	_, err = r.db.Exec(decCommentQuery, postId, userId)
	if err != nil {
		logrus.WithError(err).Error("Failed to decrement comments")
		tx.Rollback()
		return err
	}

	deleteCommentQuery := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 AND post_id = $2 AND comment_id = $3", commentsTable)
	_, err = r.db.Exec(deleteCommentQuery, userId, postId, commentId)
	if err != nil {
		logrus.WithError(err).Error("Failed to delete comment")
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
