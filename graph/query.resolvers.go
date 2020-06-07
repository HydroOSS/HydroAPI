package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/HydroOSS/HydroAPI/graph/generated"
	"github.com/HydroOSS/HydroAPI/graph/model"
	rethinkdb "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func (r *queryResolver) User(ctx context.Context, id *string) (*model.User, error) {
	cursor, err := rethinkdb.DB("Hydro").Table("users").Get(id).Run(r.DB)
	if err != nil {
		return nil, err
	}

	var user model.User
	cursor.One(&user)
	return &user, nil
}

func (r *queryResolver) Server(ctx context.Context, id *string) (*model.Server, error) {
	cursor, err := rethinkdb.DB("Hydro").Table("servers").Get(id).Run(r.DB)
	if err != nil {
		return nil, err
	}

	var server model.Server
	cursor.One(&server)
	return &server, nil
}

func (r *queryResolver) Giveaway(ctx context.Context, id *string) (*model.Giveaway, error) {
	cursor, err := rethinkdb.DB("Hydro").Table("giveaways").Get(id).Run(r.DB)
	if err != nil {
		return nil, err
	}

	var giveaway model.Giveaway
	cursor.One(&giveaway)
	return &giveaway, nil
}

func (r *queryResolver) Strike(ctx context.Context, id *string) (*model.Strike, error) {
	cursor, err := rethinkdb.DB("Hydro").Table("strikes").Get(id).Run(r.DB)
	if err != nil {
		return nil, err
	}

	var strike model.Strike
	cursor.One(&strike)
	return &strike, nil
}

func (r *queryResolver) Merit(ctx context.Context, id *string) (*model.Merit, error) {
	cursor, err := rethinkdb.DB("Hydro").Table("merits").Get(id).Run(r.DB)
	if err != nil {
		return nil, err
	}

	var merit model.Merit
	cursor.One(&merit)
	return &merit, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
