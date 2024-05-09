package usecase

import "github.com/isido5ik/StoryPublishingPlatform/dtos"

func (u *usecase) GetUsers() ([]dtos.User, error) {
	return u.repos.GetUsers()
}

func (u *usecase) GetUserById(userId int) (dtos.User, error) {
	return u.repos.GetUserById(userId)
}

func (u *usecase) DeleteUser(userId int) error {
	return u.repos.DeleteUser(userId)
}
