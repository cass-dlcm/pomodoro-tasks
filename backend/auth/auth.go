package auth

import (
	"errors"
	"github.com/cass-dlcm/pomodoro_tasks/backend/db"
	"github.com/cass-dlcm/pomodoro_tasks/graph/model"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type customClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func CreateToken(user string) (string, error) {
	claims := customClaims{
		Username: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
			Issuer:    "pomodoro-tasks",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secureSecretText"))
}

func CreateUser(user model.UserAuth) (*model.User, error) {
	if _, err := db.GetUserUsername(user.Name); errors.Is(err, errors.New("user doesn't exist")) {
		if err != nil {
			return nil, err
		}
		return nil, errors.New("user already exists")
	}
	newUserAuth := model.UserAuth{
		Name: user.Name,
		Password: hashAndSalt([]byte(user.Password)),
	}
	newUser, err := db.CreateUser(newUserAuth)
	if err != nil {
		return nil, err
	}
	list, err := db.CreateList(*newUser, "My List")
	if err != nil {
		return nil, err
	}
	newUser.Lists = []int64{*list}
	return newUser, nil
}

func CheckPassword(user model.UserAuth) error {
	userFromDb, err := db.GetUserAuthUsername(user.Name)
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userFromDb.Password), []byte(user.Password)); err != nil {
		return errors.New("incorrect password")
	}
	return nil
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
