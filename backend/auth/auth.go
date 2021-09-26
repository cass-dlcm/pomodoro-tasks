package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/cass-dlcm/pomodoro_tasks/backend/db"
	"github.com/cass-dlcm/pomodoro_tasks/graph/model"
	"github.com/form3tech-oss/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strings"
	"time"
)

const SECRETKEY = "key"

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			log.Println("Malformed token")
			next.ServeHTTP(w, r)
			return
		}
		jwtToken := authHeader[1]
		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(SECRETKEY), nil
		})
		if err != nil {
			log.Println(err)
		}
		log.Println(token.Valid)
		if token.Valid {
			ctx := context.WithValue(r.Context(), "user", token.Claims)
			// Access context values in handlers like this
			// props, _ := r.Context().Value("props").(jwt.MapClaims)
			log.Println(token.Claims.(jwt.MapClaims)["Username"])
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			log.Println("claims not okay")
			log.Println(token.Claims)
			next.ServeHTTP(w, r)
		}
	})
}

type customClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func CreateToken(user string) (string, error) {
	claims := customClaims{
		Username: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
			Issuer:    "pomodoro-tasks",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(SECRETKEY))
}

func CreateUser(user model.UserAuth) (*model.User, error) {
	if _, err := db.GetUserUsername(user.Name); errors.Is(err, errors.New("user doesn't exist")) {
		if err != nil {
			return nil, err
		}
		return nil, errors.New("user already exists")
	}
	newUserAuth := model.UserAuth{
		Name:     user.Name,
		Password: hashAndSalt([]byte(user.Password)),
	}
	log.Println(newUserAuth.Password)
	newUser := &model.User{
		Name:  user.Name,
		Lists: nil,
	}
	var err error
	newUser.ID, err = db.CreateUser(newUserAuth)
	if err != nil {
		return nil, err
	}
	list, err := db.CreateList(newUser.ID, "My List")
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

func GetUsername(ctx context.Context) string {
	if ctx.Value("user").(jwt.MapClaims)["username"] == nil {
		return ""
	}
	return ctx.Value("user").(jwt.MapClaims)["username"].(string)
}

func CheckPermsTodo(todoId int64, ctx context.Context) error {
	user, err := db.GetUserUsername(GetUsername(ctx))
	if err != nil {
		return err
	}
	todo, err := db.GetTodo(todoId)
	if err != nil {
		return err
	}
	taskList, err := db.GetListOnlyUsers(todo.List)
	if err != nil {
		return err
	}
	for i := 0; i < len(taskList.Users); i++ {
		if user.ID == taskList.Users[i] {
			return nil
		}
	}
	return errors.New("user doesn't have permission to modify this todo")
}

func CheckPermsList(listId int64, ctx context.Context) error {
	user, err := db.GetUserUsername(GetUsername(ctx))
	if err != nil {
		return err
	}
	taskList, err := db.GetListOnlyUsers(listId)
	if err != nil {
		return err
	}
	for i := 0; i < len(taskList.Users); i++ {
		if user.ID == taskList.Users[i] {
			return nil
		}
	}
	return errors.New("user doesn't have permission to modify this todo")
}
