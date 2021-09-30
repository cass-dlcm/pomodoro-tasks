package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/cass-dlcm/pomodoro_tasks/backend/application_errors"
	"github.com/cass-dlcm/pomodoro_tasks/backend/auth"
	"github.com/cass-dlcm/pomodoro_tasks/backend/db"
	"github.com/cass-dlcm/pomodoro_tasks/graph/generated"
	"github.com/cass-dlcm/pomodoro_tasks/graph/model"
)

func (r *mutationResolver) AddDependencyTodo(ctx context.Context, dependent int64, dependsOn int64) ([]*model.Todo, error) {
	if err := auth.CheckPermsTodo(dependent, ctx); err != nil {
		if errors.Is(err, application_errors.ErrNoUser) {
			log.Printf("attempt to remove dependency %d:%d by nonexistant user", dependent, dependsOn)
			return nil, application_errors.ErrPleaseAuth
		}
		if errors.Is(err, application_errors.ErrCannotFetchTodoItemNoPrint(dependent, "")) {
			return nil, application_errors.ErrCannotFetchTodoItemNoPrint(dependent, " dependent")
		}
		if errors.Is(err, application_errors.ErrNoPermissionItemNoPrint(dependent, " todo")) {
			return nil, err
		}
		return nil, application_errors.ErrUnspecified(err)
	}
	if err := auth.CheckPermsTodo(dependsOn, ctx); err != nil {
		if errors.Is(err, application_errors.ErrCannotFetchTodoItemNoPrint(dependsOn, "")) {
			return nil, application_errors.ErrCannotFetchTodoItemNoPrint(dependsOn, " dependsOn")
		}
		if errors.Is(err, application_errors.ErrNoPermissionItemNoPrint(dependsOn, " todo")) {
			return nil, err
		}
		return nil, application_errors.ErrUnspecified(err)
	}
	log.Println(GetPreloads(ctx))
	ok, err := db.CheckSameList(dependent, dependsOn)
	if err != nil {
		return nil, application_errors.ErrUnspecified(err)
	}
	if !ok {
		dependentTodo, _ := db.GetTodoStub(dependent)
		dependsOnTodo, _ := db.GetTodoStub(dependsOn)
		return nil, application_errors.ErrNotSameList(*dependentTodo, *dependsOnTodo)
	}
	found, err := db.CheckDependency(dependent, dependsOn)
	if found == true {
		return nil, application_errors.ErrDependencyFound
	}
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, application_errors.ErrUnspecified(err)
	}
	todo, err := db.AddDependency(dependent, dependsOn)
	if err != nil {
		return nil, application_errors.ErrUnspecified(err)
	}
	return todo, nil
}

func (r *mutationResolver) RemoveDependencyTodo(ctx context.Context, dependent int64, dependsOn int64) (*bool, error) {
	if err := auth.CheckPermsTodo(dependent, ctx); err != nil {
		if errors.Is(err, application_errors.ErrNoUser) {
			log.Printf("attempt to remove dependency %d:%d by nonexistant user", dependent, dependsOn)
			return nil, application_errors.ErrPleaseAuth
		}
		if errors.Is(err, application_errors.ErrCannotFetchTodoItemNoPrint(dependent, "")) {
			return nil, application_errors.ErrCannotFetchTodoItemNoPrint(dependent, " dependent")
		}
		if errors.Is(err, application_errors.ErrNoPermissionItemNoPrint(dependent, " todo")) {
			return nil, err
		}
		return nil, application_errors.ErrUnspecified(err)
	}
	if err := auth.CheckPermsTodo(dependsOn, ctx); err != nil {
		if errors.Is(err, application_errors.ErrCannotFetchTodoItemNoPrint(dependsOn, "")) {
			return nil, application_errors.ErrCannotFetchTodoItemNoPrint(dependsOn, " dependsOn")
		}
		if errors.Is(err, application_errors.ErrNoPermissionItemNoPrint(dependsOn, " todo")) {
			return nil, err
		}
		return nil, application_errors.ErrUnspecified(err)
	}
	ok, err := db.CheckSameList(dependent, dependsOn)
	if err != nil {
		return nil, application_errors.ErrUnspecified(err)
	}
	if !ok {
		dependentTodo, _ := db.GetTodoStub(dependent)
		dependsOnTodo, _ := db.GetTodoStub(dependsOn)
		return nil, application_errors.ErrNotSameList(*dependentTodo, *dependsOnTodo)
	}
	if found, err := db.CheckDependency(dependent, dependsOn); found == false {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, application_errors.ErrUnspecified(err)
		}
		return nil, application_errors.ErrNoDependency
	}
	ok, err = db.RemoveDependency(dependent, dependsOn)
	if err != nil {
		return nil, application_errors.ErrUnspecified(err)
	}
	return &ok, nil
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	if _, err := db.GetUserUsername(auth.GetUsername(ctx)); err != nil {
		if errors.Is(err, application_errors.ErrNoUser) {
			return nil, err
		}
		return nil, application_errors.ErrUnspecified(err)
	}
	log.Println(GetPreloads(ctx))
	todo := &model.Todo{
		Name:          input.Name,
		CreatedAt:     time.Now(),
		ModifiedAt:    time.Now(),
		CompletedAt:   nil,
		List:          input.List,
		DependsOnThis: nil,
		ThisDependsOn: nil,
	}
	todoId, err := db.CreateTodo(*todo)
	if err != nil {
		return nil, application_errors.ErrUnspecified(err)
	}
	todo.ID = *todoId
	return todo, nil
}

func (r *mutationResolver) RenameTodo(ctx context.Context, id int64, newName string) (*model.Todo, error) {
	if err := auth.CheckPermsTodo(id, ctx); err != nil {
		if errors.Is(err, application_errors.ErrNoPermissionItemNoPrint(id, " todo")) || errors.Is(err, application_errors.ErrCannotFetchTodoItemNoPrint(id, "")) {
			return nil, err
		}
		return nil, application_errors.ErrUnspecified(err)
	}
	log.Println(GetPreloads(ctx))
	todo, err := db.RenameTodo(id, newName)
	if err != nil {
		return nil, application_errors.ErrUnspecified(err)
	}
	return todo, nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, id int64) (*bool, error) {
	if err := auth.CheckPermsTodo(id, ctx); err != nil {
		if errors.Is(err, application_errors.ErrNoPermissionItemNoPrint(id, " todo")) || errors.Is(err, application_errors.ErrCannotFetchTodoItemNoPrint(id, "")) {
			return nil, err
		}
		return nil, application_errors.ErrUnspecified(err)
	}
	if err := db.DeleteTodo(id); err != nil {
		return nil, application_errors.ErrUnspecified(err)
	}
	trueVal := true
	return &trueVal, nil
}

func (r *mutationResolver) MarkCompletedTodo(ctx context.Context, id int64) (*model.Todo, error) {
	if err := auth.CheckPermsTodo(id, ctx); err != nil {
		if errors.Is(err, application_errors.ErrNoPermissionItemNoPrint(id, " todo")) || errors.Is(err, application_errors.ErrCannotFetchTodoItemNoPrint(id, "")) {
			return nil, err
		}
		return nil, application_errors.ErrUnspecified(err)
	}
	log.Println(GetPreloads(ctx))
	todo, err := db.UpdateCompletionTodo(id)
	if err != nil {
		return nil, application_errors.ErrUnspecified(err)
	}
	return todo, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, user model.UserAuth) (*model.User, error) {
	resultUser, err := auth.CreateUser(user)
	if err != nil {
		if errors.Is(err, application_errors.ErrUserExists) {
			return nil, err
		}
		return nil, application_errors.ErrUnspecified(err)
	}
	return resultUser, nil
}

func (r *mutationResolver) SignIn(ctx context.Context, user model.UserAuth) (*string, error) {
	if err := auth.CheckPassword(user); err != nil {
		if errors.Is(err, application_errors.ErrNoUser) || errors.Is(err, application_errors.ErrIncorrectPass) {
			return nil, err
		}
		return nil, application_errors.ErrUnspecified(err)
	}
	token, err := auth.CreateToken(user.Name)
	if err != nil {
		return nil, application_errors.ErrUnspecified(err)
	}
	return &token, nil
}

func (r *queryResolver) Todos(ctx context.Context, list int64) (*model.TaskList, error) {
	if err := auth.CheckPermsList(list, ctx); err != nil {
		if errors.Is(err, application_errors.ErrNoPermissionItemNoPrint(list, " list")) || errors.Is(err, application_errors.ErrCannotFetchTodoListNoPrint(list)) || errors.Is(err, application_errors.ErrNoUser) {
			return nil, err
		}
		return nil, application_errors.ErrUnspecified(err)
	}
	log.Println(GetPreloads(ctx))
	return db.GetListOnlyTasks(list)
}

func (r *queryResolver) Lists(ctx context.Context) ([]int64, error) {
	username := auth.GetUsername(ctx)
	if username == "" {
		return nil, errors.New("user was not found")
	}
	user, err := db.GetUserUsername(username)
	if err != nil {
		if errors.Is(err, application_errors.ErrNoUser) {
			return nil, err
		}
		return nil, application_errors.ErrUnspecified(err)
	}
	log.Println(GetPreloads(ctx))
	taskLists, err := db.GetTaskListsUser(user.ID)
	if err != nil {
		return nil, application_errors.ErrUnspecified(err)
	}
	return taskLists, nil
}

func (r *queryResolver) GetTodo(ctx context.Context, id int64) (*model.Todo, error) {
	if err := auth.CheckPermsTodo(id, ctx); err != nil {
		if errors.Is(err, application_errors.ErrNoPermissionItemNoPrint(id, " todo")) {
			return nil, err
		}
		if errors.Is(err, application_errors.ErrCannotFetchTodoItemNoPrint(id, "")) {
			return nil, err
		}
		return nil, application_errors.ErrUnspecified(err)
	}
	log.Println(GetPreloads(ctx))
	return db.GetTodo(id)
}

func (r *queryResolver) CheckDependencyTodo(ctx context.Context, dependent int64, dependsOn int64) (*bool, error) {
	if err := auth.CheckPermsTodo(dependent, ctx); err != nil {
		if errors.Is(err, application_errors.ErrNoPermissionItemNoPrint(dependent, " todo")) {
			return nil, err
		}
		if errors.Is(err, application_errors.ErrCannotFetchTodoItemNoPrint(dependent, "")) {
			return nil, application_errors.ErrCannotFetchTodoItem(dependent, " dependent")
		}
		return nil, application_errors.ErrUnspecified(err)
	}
	if err := auth.CheckPermsTodo(dependsOn, ctx); err != nil {
		if errors.Is(err, application_errors.ErrNoPermissionItemNoPrint(dependsOn, " todo")) {
			return nil, err
		}
		if errors.Is(err, application_errors.ErrCannotFetchTodoItemNoPrint(dependsOn, "")) {
			return nil, application_errors.ErrCannotFetchTodoItem(dependsOn, " dependsOn")
		}
		return nil, application_errors.ErrUnspecified(err)
	}
	found, err := db.CheckDependency(dependent, dependsOn)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return &found, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
