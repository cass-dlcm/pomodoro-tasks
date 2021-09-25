package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/cass-dlcm/pomodoro_tasks/backend/db"
	"time"

	"github.com/cass-dlcm/pomodoro_tasks/backend/auth"
	"github.com/cass-dlcm/pomodoro_tasks/graph/generated"
	"github.com/cass-dlcm/pomodoro_tasks/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo, list string) (*model.Todo, error) {
	todo := &model.Todo{
		Name:        input.Name,
		CreatedAt:   time.Now(),
		ModifiedAt:  time.Now(),
		CompletedAt: nil,
		List:        input.List,
	}
	return todo, nil
}

func (r *mutationResolver) RenameTodo(ctx context.Context, id int64, newName string) (*model.Todo, error) {
	todo, err := db.GetTodo(id)
	if err != nil {
		return nil, err
	}
	todo.Name = newName
	db.EditTodo(todo)
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input int64) (*bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) MarkCompletedTodo(ctx context.Context, input int64) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, user model.UserAuth) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SignIn(ctx context.Context, user model.UserAuth) (*string, error) {
	if err := auth.CheckPassword(user); err != nil {
		return nil, err
	}
	token, err := auth.CreateToken(user.Name)
	return &token, err
}

func (r *queryResolver) Todos(ctx context.Context, list int64) ([]int64, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Lists(ctx context.Context) ([]int64, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]int64, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
