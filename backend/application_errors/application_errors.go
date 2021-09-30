package application_errors

import (
	"errors"
	"fmt"
	"github.com/cass-dlcm/pomodoro_tasks/graph/model"
	"log"
	"time"
)

var ErrNoUser = errors.New("no user found")
var ErrUserExists = errors.New("user already exists")
var ErrIncorrectPass = errors.New("incorrect password")
var ErrPleaseAuth = errors.New("please log in and try again")
var ErrNoDependency = errors.New("no dependency found")
var ErrDependencyFound = errors.New("dependency found")

func ErrCannotFetchTodoItem(item int64, relation string) error {
	log.Printf(ErrCannotFetchTodoItemNoPrint(item, relation).Error())
	return ErrCannotFetchTodoItemNoPrint(item, relation)
}

func ErrCannotFetchTodoItemNoPrint(item int64, relation string) error {
	return errors.New(fmt.Sprintf("cannot fetch%s todo item: %d", relation, item))
}

func ErrNoPermissionItem(item int64, itemtype, user string) error {
	log.Printf("permission denied for access of%s %d by user %s", itemtype, item, user)
	return ErrNoPermissionItemNoPrint(item, itemtype)
}

func ErrNoPermissionItemNoPrint(item int64, itemtype string) error {
	return errors.New(fmt.Sprintf("you dont have permission to access or modify the%s item: %d", itemtype, item))
}

func ErrNotSameList(dependent model.TodoStub, dependsOn model.TodoStub) error {
	return errors.New(fmt.Sprintf("the todo items are not on the same list\ndependent item %d is on list %d, dependsOn item %d is on list %d", dependent.ID, dependent.List, dependsOn.ID, dependsOn.List))
}

func ErrCannotFetchTodoList(item int64) error {
	log.Println(ErrCannotFetchTodoListNoPrint(item).Error())
	return ErrCannotFetchTodoListNoPrint(item)
}

func ErrCannotFetchTodoListNoPrint(item int64) error {
	return errors.New(fmt.Sprintf("cannot fetch todo list: %d", item))
}

func ErrUnspecified(err error) error {
	log.Println(err)
	return errors.New(fmt.Sprintf("%v: an unspecified error has occured\nplease let the admin know and give this timestamp", time.Now().Format("2006-01-02 15:04:05")))
}
