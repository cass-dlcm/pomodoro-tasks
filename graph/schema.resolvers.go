package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"math"
	"time"

	"github.com/cass-dlcm/pomodoro_tasks/backend/application_errors"
	"github.com/cass-dlcm/pomodoro_tasks/backend/auth"
	"github.com/cass-dlcm/pomodoro_tasks/backend/db"
	"github.com/cass-dlcm/pomodoro_tasks/graph/generated"
	"github.com/cass-dlcm/pomodoro_tasks/graph/model"
)

func (*mutationResolver) AddDependencyTodo(ctx context.Context, dependent int64, dependsOn int64) ([]*model.Todo, error) {
	if err := auth.CheckPermsTodo(ctx, dependent); err != nil {
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
	if err := auth.CheckPermsTodo(ctx, dependsOn); err != nil {
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
	if found {
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

func (*mutationResolver) RemoveDependencyTodo(ctx context.Context, dependent int64, dependsOn int64) (*bool, error) {
	if err := auth.CheckPermsTodo(ctx, dependent); err != nil {
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
	if err := auth.CheckPermsTodo(ctx, dependsOn); err != nil {
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
	if found, err := db.CheckDependency(dependent, dependsOn); !found {
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

func (*mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
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

func (*mutationResolver) RenameTodo(ctx context.Context, id int64, newName string) (*model.Todo, error) {
	if err := auth.CheckPermsTodo(ctx, id); err != nil {
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

func (*mutationResolver) DeleteTodo(ctx context.Context, id int64) (*bool, error) {
	if err := auth.CheckPermsTodo(ctx, id); err != nil {
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

func (*mutationResolver) MarkCompletedTodo(ctx context.Context, id int64) (*model.Todo, error) {
	if err := auth.CheckPermsTodo(ctx, id); err != nil {
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

func (*mutationResolver) CreateUser(ctx context.Context, user model.UserAuth) (*model.User, error) {
	if auth.GetUsername(ctx) != "" {
		return nil, application_errors.ErrAlreadySignedIn
	}
	resultUser, err := auth.CreateUser(user)
	if err != nil {
		if errors.Is(err, application_errors.ErrUserExists) {
			return nil, err
		}
		return nil, application_errors.ErrUnspecified(err)
	}
	return resultUser, nil
}

func (*mutationResolver) SignIn(ctx context.Context, user model.UserAuth) (*string, error) {
	if auth.GetUsername(ctx) != "" {
		return nil, application_errors.ErrAlreadySignedIn
	}
	userInfo, err := db.GetUserUsername(user.Name)
	if err != nil {
		if errors.Is(err, application_errors.ErrNoUser) {
			return nil, err
		}
		return nil, application_errors.ErrUnspecified(err)
	}
	count, lastFailedLogin, err := db.GetTimeout(userInfo.ID, ctx.Value(auth.ContextKey("ip")).(string))
	if err == nil {
		if time.Now().Before(lastFailedLogin.Add(time.Second * 1 << *count)) {
			return nil, application_errors.ErrPleaseWaitForAuth(user.Name, int64(math.Ceil(lastFailedLogin.Add(time.Second*2<<*count).Sub(time.Now()).Seconds())))
		}
	}
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, application_errors.ErrUnspecified(err)
	}
	if count == nil {
		countVal := 0
		count = &countVal
	}
	if err := auth.CheckPassword(user, count); err != nil {
		if err.Error() == application_errors.ErrIncorrectPassNoPrint(user.Name, 1<<(*count+1)).Error() {
			if err := db.IncrementTimeout(userInfo.ID, ctx.Value(auth.ContextKey("ip")).(string), *count+1); err != nil {
				return nil, application_errors.ErrUnspecified(err)
			}
			return nil, err
		}
		return nil, application_errors.ErrUnspecified(err)
	}
	if err := db.DeleteTimeout(userInfo.ID, ctx.Value(auth.ContextKey("ip")).(string)); err != nil {
		return nil, application_errors.ErrUnspecified(err)
	}
	token, err := auth.CreateToken(user.Name)
	if err != nil {
		return nil, application_errors.ErrUnspecified(err)
	}
	return &token, nil
}

func (*queryResolver) Todos(ctx context.Context, list int64) (*model.TaskList, error) {
	if err := auth.CheckPermsList(ctx, list); err != nil {
		if errors.Is(err, application_errors.ErrNoPermissionItemNoPrint(list, " list")) || errors.Is(err, application_errors.ErrCannotFetchTodoListNoPrint(list)) || errors.Is(err, application_errors.ErrNoUser) {
			return nil, err
		}
		return nil, application_errors.ErrUnspecified(err)
	}
	log.Println(GetPreloads(ctx))
	return db.GetListOnlyTasks(list)
}

func (*queryResolver) Lists(ctx context.Context) ([]int64, error) {
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

func (*queryResolver) GetTodo(ctx context.Context, id int64) (*model.Todo, error) {
	if err := auth.CheckPermsTodo(ctx, id); err != nil {
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

func (*queryResolver) CheckDependencyTodo(ctx context.Context, dependent int64, dependsOn int64) (*bool, error) {
	if err := auth.CheckPermsTodo(ctx, dependent); err != nil {
		if errors.Is(err, application_errors.ErrNoPermissionItemNoPrint(dependent, " todo")) {
			return nil, err
		}
		if errors.Is(err, application_errors.ErrCannotFetchTodoItemNoPrint(dependent, "")) {
			return nil, application_errors.ErrCannotFetchTodoItem(dependent, " dependent")
		}
		return nil, application_errors.ErrUnspecified(err)
	}
	if err := auth.CheckPermsTodo(ctx, dependsOn); err != nil {
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
