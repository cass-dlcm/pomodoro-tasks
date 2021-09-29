package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/cass-dlcm/pomodoro_tasks/backend/auth"
	"github.com/cass-dlcm/pomodoro_tasks/backend/db"
	"github.com/cass-dlcm/pomodoro_tasks/graph/generated"
	"github.com/cass-dlcm/pomodoro_tasks/graph/model"
)

func (r *mutationResolver) AddDependencyTodo(ctx context.Context, dependent int64, dependsOn int64) ([]*model.Todo, error) {
	if err := auth.CheckPermsTodo(dependent, ctx); err != nil {
		return nil, err
	}
	if err := auth.CheckPermsTodo(dependsOn, ctx); err != nil {
		return nil, err
	}
	if ok, err := db.CheckSameList(dependent, dependsOn); err != nil || !ok {
		return nil, err
	}
	if found, err := db.CheckDependency(dependent, dependsOn); found == true {
		return nil, err
	}
	todo, err := db.AddDependency(dependent, dependsOn)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *mutationResolver) RemoveDependencyTodo(ctx context.Context, dependent int64, dependsOn int64) (bool, error) {
	if err := auth.CheckPermsTodo(dependent, ctx); err != nil {
		return false, err
	}
	if err := auth.CheckPermsTodo(dependsOn, ctx); err != nil {
		return false, err
	}
	if ok, err := db.CheckSameList(dependent, dependsOn); err != nil || !ok {
		return true, err
	}
	if found, err := db.CheckDependency(dependent, dependsOn); found == false {
		if !errors.Is(err, sql.ErrNoRows) {
			return true, err
		}
		return true, nil
	}
	return db.RemoveDependency(dependent, dependsOn)
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	if _, err := db.GetUserUsername(auth.GetUsername(ctx)); err != nil {
		return nil, err
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
		return nil, err
	}
	todo.ID = *todoId
	return todo, nil
}

func (r *mutationResolver) RenameTodo(ctx context.Context, id int64, newName string) (*model.Todo, error) {
	if err := auth.CheckPermsTodo(id, ctx); err != nil {
		return nil, err
	}
	log.Println(GetPreloads(ctx))
	return db.RenameTodo(id, newName)
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, id int64) (bool, error) {
	if err := auth.CheckPermsTodo(id, ctx); err != nil {
		return false, err
	}
	if err := db.DeleteTodo(id); err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) MarkCompletedTodo(ctx context.Context, id int64) (*model.Todo, error) {
	if err := auth.CheckPermsTodo(id, ctx); err != nil {
		return nil, err
	}
	log.Println(GetPreloads(ctx))
	return db.UpdateCompletionTodo(id)
}

func (r *mutationResolver) CreateUser(ctx context.Context, user model.UserAuth) (*model.User, error) {
	return auth.CreateUser(user)
}

func (r *mutationResolver) SignIn(ctx context.Context, user model.UserAuth) (*string, error) {
	if err := auth.CheckPassword(user); err != nil {
		return nil, err
	}
	token, err := auth.CreateToken(user.Name)
	return &token, err
}

func (r *queryResolver) Todos(ctx context.Context, list int64) (*model.TaskList, error) {
	if err := auth.CheckPermsList(list, ctx); err != nil {
		return nil, err
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
		return nil, err
	}
	log.Println(GetPreloads(ctx))
	return db.GetTaskListsUser(user.ID)
}

func (r *queryResolver) GetTodo(ctx context.Context, id int64) (*model.Todo, error) {
	if err := auth.CheckPermsTodo(id, ctx); err != nil {
		return nil, err
	}
	return db.GetTodo(id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

