package usecase

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/isido5ik/StoryPublishingPlatform/dtos"
	"github.com/sirupsen/logrus"
)

const (
	signingKey = "qrkjk#4#5FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int          `json:"userId"`
	Roles  []dtos.Roles `json:"roles"`
}

func (u *usecase) CreateUserAsClient(input dtos.SignUpInput) (int, error) {
	_, err := u.repos.GetUser(input.Username, generateHashPassword(input.Password))
	if err == nil {
		return 0, errors.New("user with this username already exists")
	}

	input.Password = generateHashPassword(input.Password)
	return u.repos.CreateUserAsClient(input)
}

func (u *usecase) GenerateToken(username, password string) (string, []dtos.Roles, error) {
	user, err := u.repos.GetUser(username, generateHashPassword(password))
	if err != nil {
		return "", nil, err
	}

	logrus.WithField("user", user.Username).Info("\nGetting user to generate token")

	var rolesHeaders []dtos.Roles
	roles, err := u.repos.GetRoles(user.UserID)
	log.Println(roles)
	if err != nil {

		return "", nil, err
	}

	for _, role := range roles {
		role_id, err := u.repos.GetRoleId(role, user.UserID)
		logrus.WithField("role_id", role_id)
		logrus.WithField("role", role).Info("Getting role id\n")

		if err != nil {
			logrus.WithError(err).Error("Failed to get role id")
			return "", nil, err
		}
		rolesHeaders = append(rolesHeaders, dtos.Roles{RoleId: role_id, RoleName: role})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.UserID,
		rolesHeaders,
	})

	tokenString, err := token.SignedString([]byte(signingKey))
	if err != nil {
		logrus.WithError(err).Error("Error signing string token")
		return "", nil, err
	}

	return tokenString, rolesHeaders, nil
}

func (u *usecase) ParseToken(tokenString string) (int, []dtos.Roles, error) {

	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		logrus.WithError(err).Error("Failed to parse token")
		return 0, nil, err
	}

	if !token.Valid {
		return 0, nil, errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, nil, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, claims.Roles, nil
}

func generateHashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum(nil))
}
