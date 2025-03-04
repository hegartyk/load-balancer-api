package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.30

import (
	"context"

	"go.infratographer.com/load-balancer-api/internal/ent/generated"
	"go.infratographer.com/x/gidx"
)

// LoadBalancerOriginCreate is the resolver for the loadBalancerOriginCreate field.
func (r *mutationResolver) LoadBalancerOriginCreate(ctx context.Context, input generated.CreateLoadBalancerOriginInput) (*LoadBalancerOriginCreatePayload, error) {
	// check if pool exists
	_, err := r.client.Pool.Get(ctx, input.PoolID)
	if err != nil {
		return nil, err
	}

	origin, err := r.client.Origin.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &LoadBalancerOriginCreatePayload{LoadBalancerOrigin: origin}, nil
}

// LoadBalancerOriginUpdate is the resolver for the loadBalancerOriginUpdate field.
func (r *mutationResolver) LoadBalancerOriginUpdate(ctx context.Context, id gidx.PrefixedID, input generated.UpdateLoadBalancerOriginInput) (*LoadBalancerOriginUpdatePayload, error) {
	origin, err := r.client.Origin.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	origin, err = origin.Update().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &LoadBalancerOriginUpdatePayload{LoadBalancerOrigin: origin}, nil
}

// LoadBalancerOriginDelete is the resolver for the loadBalancerOriginDelete field.
func (r *mutationResolver) LoadBalancerOriginDelete(ctx context.Context, id gidx.PrefixedID) (*LoadBalancerOriginDeletePayload, error) {
	if err := r.client.Origin.DeleteOneID(id).Exec(ctx); err != nil {
		return nil, err
	}

	return &LoadBalancerOriginDeletePayload{DeletedID: id}, nil
}
