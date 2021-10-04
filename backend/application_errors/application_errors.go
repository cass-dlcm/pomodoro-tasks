package application_errors

import (
	"errors"
	"fmt"
	"github.com/cass-dlcm/pomodoro_tasks/graph/model"
	"log"
	"runtime/debug"
	"time"
)

var ErrUserExists = errors.New("user already exists")
var ErrPleaseAuth = errors.New("please log in and try again")
var ErrNoDependency = errors.New("no dependency found")
var ErrDependencyFound = errors.New("dependency found")
var ErrAlreadySignedIn = errors.New("user already signed in")

func ErrCannotFetchTodoItem(item int64, relation string) error {
	log.Println(ErrCannotFetchTodoItemNoPrint(item, relation).Error())
	return ErrCannotFetchTodoItemNoPrint(item, relation)
}

func ErrCannotFetchTodoItemNoPrint(item int64, relation string) error {
	return fmt.Errorf("cannot fetch%s todo item: %d", relation, item)
}

func ErrNoPermissionItem(item int64, itemtype, user string) error {
	log.Printf("permission denied for access of%s %d by user %s", itemtype, item, user)
	return ErrNoPermissionItemNoPrint(item, itemtype)
}

func ErrNoPermissionItemNoPrint(item int64, itemtype string) error {
	return fmt.Errorf("you dont have permission to access or modify the%s item: %d", itemtype, item)
}

func ErrNotSameList(dependent model.TodoStub, dependsOn model.TodoStub) error {
	return fmt.Errorf("the todo items are not on the same list\ndependent item %d is on list %d, dependsOn item %d is on list %d", dependent.ID, dependent.List, dependsOn.ID, dependsOn.List)
}

func ErrCannotFetchTodoList(item int64) error {
	log.Println(ErrCannotFetchTodoListNoPrint(item).Error())
	return ErrCannotFetchTodoListNoPrint(item)
}

func ErrCannotFetchTodoListNoPrint(item int64) error {
	return fmt.Errorf("cannot fetch todo list: %d", item)
}

func ErrUnspecified(err error) error {
	log.Println(err)
	debug.PrintStack()
	return fmt.Errorf("%v: an unspecified error has occured\nplease let the admin know and give this timestamp", time.Now().Format("2006-01-02 15:04:05"))
}

func ErrIncorrectPass(username string, timeout int64) error {
	log.Println(ErrIncorrectPassNoPrint(username, timeout).Error())
	return ErrIncorrectPassNoPrint(username, timeout)
}

func ErrIncorrectPassNoPrint(username string, timeout int64) error {
	return fmt.Errorf("incorrect password for user %s, please try again in %d seconds", username, timeout)
}

func ErrPleaseWaitForAuth(username string, timeout int64) error {
	log.Println(ErrPleaseWaitForAuthNoPrint(username, timeout).Error())
	return ErrPleaseWaitForAuthNoPrint(username, timeout)
}

func ErrPleaseWaitForAuthNoPrint(username string, timeout int64) error {
	return fmt.Errorf("user %s is on login timeout, please try again in %d seconds", username, timeout)
}

func ErrNoUserNoPrint(username string) error {
	return fmt.Errorf("no user found with name %s", username)
}

func ErrNoUser(username string) error {
	log.Println(ErrNoUserNoPrint(username))
	return fmt.Errorf("no user found with name %s", username)
}