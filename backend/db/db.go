package db

import (
	"database/sql"
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/cass-dlcm/pomodoro_tasks/graph/model"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("mysql", "user:password@/dbname?parseTime=true")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

}

func GetUserUsername(username string) (*model.User, error) {
	user := &model.User{}
	if err := db.QueryRow("select id, username from users where username = ?", username).Scan(&user.ID, &user.Name); err != nil {
		return nil, err
	}
	var err error
	user.Lists, err = GetTaskListsUser(user.ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	return user, nil
}

func GetUserAuthUsername(username string) (*model.UserAuth, error) {
	user := &model.UserAuth{}
	if err := db.QueryRow("select username, password from users where username = ?", username).Scan(&user.Name, &user.Password); err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(user model.UserAuth) (*model.User, error) {
	res, err := db.Exec("insert into users (username, password) values (?, ?)", user.Name, user.Password)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return GetUser(id)
}

func GetUser(id int64) (*model.User, error) {
	user := &model.User{
		ID:    0,
		Name:  "",
		Lists: []int64{},
	}
	if err := db.QueryRow("select id, username from users where id = ?", id).Scan(&user.ID, &user.Name); err != nil {
		return nil, err
	}
	var err error
	user.Lists, err = GetTaskListsUser(id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	return user, nil
}

func GetTaskListsUser(id int64) ([]int64, error) {
	taskLists := []int64{}
	rows, err := db.Query("select id from tasklist_user_link where user = ?", id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	for rows.Next() {
		var rowId int64
		if err := rows.Scan(&rowId); err != nil {
			return nil, err
		}
		taskLists = append(taskLists, rowId)
	}
	return taskLists, nil
}

func CreateList(user model.User, name string) (*int64, error) {
	res, err := db.Exec("insert into lists (listname) values (?)", name)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	res, err = db.Exec("insert into tasklist_user_link (user, list) values (?, ?)", user.ID, id)
	id, err = res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func GetTodo(id int64) (*model.Todo, error) {
	todo := model.Todo{
		ID:          id,
	}
	if err := db.QueryRow("select taskname, createdat, modifiedat, completedat, list from todos where id = ?", id).Scan(&todo.Name, &todo.CreatedAt, &todo.ModifiedAt, &todo.CompletedAt, &todo.List); err != nil {
		return nil, err
	}
	return &todo, nil
}

func GetListOnlyUsers(listId int64) (*model.TaskList, error) {
	taskList := &model.TaskList{
		ID:    listId,
		Users: []int64{},
	}
	if err := db.QueryRow("select listname from lists where id = ?", listId).Scan(&taskList.Name); err != nil {
		return nil, err
	}
	rows, err := db.Query("select user from tasklist_user_link where id = ?", listId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var userid int64
		if err := rows.Scan(&userid); err != nil {
			return nil, err
		}
		taskList.Users = append(taskList.Users, userid)
	}
	return taskList, nil
}

func GetListOnlyTasks(listId int64) (*model.TaskList, error) {
	taskList := &model.TaskList{
		ID:    listId,
		Tasks: []*model.Todo{},
	}
	if err := db.QueryRow("select listname from lists where id = ?", listId).Scan(&taskList.Name); err != nil {
		return nil, err
	}
	rows, err := db.Query("select id, taskname, createdat, modifiedat, completedat from todos where list = ?", listId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	for rows.Next() {
		todo := &model.Todo{
			List: listId,
		}
		if err := rows.Scan(&todo.ID, &todo.Name, &todo.CreatedAt, &todo.ModifiedAt, &todo.CompletedAt); err != nil {
			return nil, err
		}
		taskList.Tasks = append(taskList.Tasks, todo)
	}
	return taskList, nil
}

func RenameTodo(id int64, name string) (*model.Todo, error) {
	_, err := db.Exec("update todos set taskname = ?, modifiedat = ? where id = ?", name, time.Now(), id)
	if err != nil {
		return nil, err
	}
	return GetTodo(id)
}

func CreateTodo(todo model.Todo) (*int64, error) {
	res, err := db.Exec("insert into todos (taskname, createdat, modifiedat, completedat, list) values (?, ?, ?, ?, ?)", todo.Name, todo.CreatedAt, todo.ModifiedAt, todo.CompletedAt, todo.List)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func DeleteTodo(id int64) error {
	_, err := db.Exec("delete from todos where id = ?", id)
	return err
}

func UpdateCompletionTodo(id int64) (*model.Todo, error) {
	_, err := db.Exec("update todos set modifiedat = ?, completedat = ? where id = ?", time.Now(), time.Now(), id)
	if err != nil {
		return nil, err
	}
	return GetTodo(id)
}