package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/HydroOSS/HydroAPI/graph/generated"
	"github.com/HydroOSS/HydroAPI/graph/model"
)

func (r *queryResolver) User(ctx context.Context, id *string) (*model.User, error) {
	return &model.User{}, nil
}

func (r *queryResolver) Server(ctx context.Context, id *string) (*model.Server, error) {
	return &model.Server{}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
