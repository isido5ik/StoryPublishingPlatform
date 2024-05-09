package repository

import (
	"fmt"

	"github.com/isido5ik/StoryPublishingPlatform/dtos"
)

func (r *repository) GetUsers() ([]dtos.User, error) {
	var users []dtos.User

	getUsersQuery := fmt.Sprintf("SELECT * FROM %s", usersTable)
	err := r.db.Select(&users, getUsersQuery)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *repository) GetUserById(userId int) (dtos.User, error) {
	var user dtos.User

	getUserQuery := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", usersTable)
	if err := r.db.Get(&user, getUserQuery, userId); err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) DeleteUser(userId int) error {
	deleteUserQuery := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1", usersTable)
	_, err := r.db.Exec(deleteUserQuery, userId)
	if err != nil {
		return err
	}
	return nil
}
