package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/tmc/pulid"
	"github.com/tmc/srvgql/ent"
)

func (r *accountEdgeResolver) Cursor(ctx context.Context, obj *ent.AccountEdge) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *organizationResolver) Accounts(ctx context.Context, obj *ent.Organization) (*ent.AccountConnection, error) {
	return obj.QueryAccounts().Paginate(ctx, nil, nil, nil, nil)
}

func (r *organizationEdgeResolver) Cursor(ctx context.Context, obj *ent.OrganizationEdge) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Node(ctx context.Context, id pulid.ID) (ent.Noder, error) {
	return r.Client.Noder(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []*pulid.ID) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Me(ctx context.Context) (*ent.Account, error) {
	accounts, err := r.Client.Account.Query().All(ctx)

	fmt.Println("acc, err:", accounts, err)
	if err != nil {
		return nil, err
	}
	// TODO: this is just returning the first account
	if len(accounts) > 0 {
		return accounts[0], nil
	}
	return nil, fmt.Errorf("no accounts")
}

func (r *queryResolver) Org(ctx context.Context) (*ent.Organization, error) {
	results, err := r.Client.Organization.Query().All(ctx)
	fmt.Println("results, err:", results, err)
	if err != nil {
		return nil, err
	}
	if len(results) > 0 {
		return results[0], nil
	}
	return nil, fmt.Errorf("no results")
}

// AccountEdge returns AccountEdgeResolver implementation.
func (r *Resolver) AccountEdge() AccountEdgeResolver { return &accountEdgeResolver{r} }

// Organization returns OrganizationResolver implementation.
func (r *Resolver) Organization() OrganizationResolver { return &organizationResolver{r} }

// OrganizationEdge returns OrganizationEdgeResolver implementation.
func (r *Resolver) OrganizationEdge() OrganizationEdgeResolver { return &organizationEdgeResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type accountEdgeResolver struct{ *Resolver }
type organizationResolver struct{ *Resolver }
type organizationEdgeResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
