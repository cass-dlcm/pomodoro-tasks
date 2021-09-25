package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/cass-dlcm/pomodoro_tasks/backend/auth"
	"github.com/cass-dlcm/pomodoro_tasks/backend/db"
	"github.com/cass-dlcm/pomodoro_tasks/graph/generated"
	"github.com/cass-dlcm/pomodoro_tasks/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo, list string) (*model.Todo, error) {
	if _, err := db.GetUserUsername(auth.GetUsername(ctx)); err != nil {
		return nil, err
	}
	todo := &model.Todo{
		Name:        input.Name,
		CreatedAt:   time.Now(),
		ModifiedAt:  time.Now(),
		CompletedAt: nil,
		List:        input.List,
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
	return db.RenameTodo(id, newName)
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input int64) (bool, error) {
	if err := auth.CheckPermsTodo(input, ctx); err != nil {
		return false, err
	}
	if err := db.DeleteTodo(input); err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) MarkCompletedTodo(ctx context.Context, input int64) (*model.Todo, error) {
	if err := auth.CheckPermsTodo(input, ctx); err != nil {
		return nil, err
	}
	return db.UpdateCompletionTodo(input)
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
	return db.GetListOnlyTasks(list)
}

func (r *queryResolver) Lists(ctx context.Context) ([]int64, error) {
	user, err := db.GetUserUsername(auth.GetUsername(ctx))
	if err != nil {
		 return nil, err
	}
	return db.GetTaskListsUser(user.ID)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
