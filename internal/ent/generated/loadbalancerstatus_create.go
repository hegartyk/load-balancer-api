// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/loadbalancer"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/loadbalancerstatus"
	"go.infratographer.com/x/gidx"
)

// LoadBalancerStatusCreate is the builder for creating a LoadBalancerStatus entity.
type LoadBalancerStatusCreate struct {
	config
	mutation *LoadBalancerStatusMutation
	hooks    []Hook
}

// SetNamespace sets the "namespace" field.
func (lbsc *LoadBalancerStatusCreate) SetNamespace(s string) *LoadBalancerStatusCreate {
	lbsc.mutation.SetNamespace(s)
	return lbsc
}

// SetData sets the "data" field.
func (lbsc *LoadBalancerStatusCreate) SetData(jm json.RawMessage) *LoadBalancerStatusCreate {
	lbsc.mutation.SetData(jm)
	return lbsc
}

// SetCreatedAt sets the "created_at" field.
func (lbsc *LoadBalancerStatusCreate) SetCreatedAt(t time.Time) *LoadBalancerStatusCreate {
	lbsc.mutation.SetCreatedAt(t)
	return lbsc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lbsc *LoadBalancerStatusCreate) SetNillableCreatedAt(t *time.Time) *LoadBalancerStatusCreate {
	if t != nil {
		lbsc.SetCreatedAt(*t)
	}
	return lbsc
}

// SetUpdatedAt sets the "updated_at" field.
func (lbsc *LoadBalancerStatusCreate) SetUpdatedAt(t time.Time) *LoadBalancerStatusCreate {
	lbsc.mutation.SetUpdatedAt(t)
	return lbsc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lbsc *LoadBalancerStatusCreate) SetNillableUpdatedAt(t *time.Time) *LoadBalancerStatusCreate {
	if t != nil {
		lbsc.SetUpdatedAt(*t)
	}
	return lbsc
}

// SetLoadBalancerID sets the "load_balancer_id" field.
func (lbsc *LoadBalancerStatusCreate) SetLoadBalancerID(gi gidx.PrefixedID) *LoadBalancerStatusCreate {
	lbsc.mutation.SetLoadBalancerID(gi)
	return lbsc
}

// SetSource sets the "source" field.
func (lbsc *LoadBalancerStatusCreate) SetSource(s string) *LoadBalancerStatusCreate {
	lbsc.mutation.SetSource(s)
	return lbsc
}

// SetID sets the "id" field.
func (lbsc *LoadBalancerStatusCreate) SetID(gi gidx.PrefixedID) *LoadBalancerStatusCreate {
	lbsc.mutation.SetID(gi)
	return lbsc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (lbsc *LoadBalancerStatusCreate) SetNillableID(gi *gidx.PrefixedID) *LoadBalancerStatusCreate {
	if gi != nil {
		lbsc.SetID(*gi)
	}
	return lbsc
}

// SetLoadBalancer sets the "load_balancer" edge to the LoadBalancer entity.
func (lbsc *LoadBalancerStatusCreate) SetLoadBalancer(l *LoadBalancer) *LoadBalancerStatusCreate {
	return lbsc.SetLoadBalancerID(l.ID)
}

// Mutation returns the LoadBalancerStatusMutation object of the builder.
func (lbsc *LoadBalancerStatusCreate) Mutation() *LoadBalancerStatusMutation {
	return lbsc.mutation
}

// Save creates the LoadBalancerStatus in the database.
func (lbsc *LoadBalancerStatusCreate) Save(ctx context.Context) (*LoadBalancerStatus, error) {
	lbsc.defaults()
	return withHooks[*LoadBalancerStatus, LoadBalancerStatusMutation](ctx, lbsc.sqlSave, lbsc.mutation, lbsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lbsc *LoadBalancerStatusCreate) SaveX(ctx context.Context) *LoadBalancerStatus {
	v, err := lbsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lbsc *LoadBalancerStatusCreate) Exec(ctx context.Context) error {
	_, err := lbsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lbsc *LoadBalancerStatusCreate) ExecX(ctx context.Context) {
	if err := lbsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lbsc *LoadBalancerStatusCreate) defaults() {
	if _, ok := lbsc.mutation.CreatedAt(); !ok {
		v := loadbalancerstatus.DefaultCreatedAt()
		lbsc.mutation.SetCreatedAt(v)
	}
	if _, ok := lbsc.mutation.UpdatedAt(); !ok {
		v := loadbalancerstatus.DefaultUpdatedAt()
		lbsc.mutation.SetUpdatedAt(v)
	}
	if _, ok := lbsc.mutation.ID(); !ok {
		v := loadbalancerstatus.DefaultID()
		lbsc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lbsc *LoadBalancerStatusCreate) check() error {
	if _, ok := lbsc.mutation.Namespace(); !ok {
		return &ValidationError{Name: "namespace", err: errors.New(`generated: missing required field "LoadBalancerStatus.namespace"`)}
	}
	if v, ok := lbsc.mutation.Namespace(); ok {
		if err := loadbalancerstatus.NamespaceValidator(v); err != nil {
			return &ValidationError{Name: "namespace", err: fmt.Errorf(`generated: validator failed for field "LoadBalancerStatus.namespace": %w`, err)}
		}
	}
	if _, ok := lbsc.mutation.Data(); !ok {
		return &ValidationError{Name: "data", err: errors.New(`generated: missing required field "LoadBalancerStatus.data"`)}
	}
	if _, ok := lbsc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generated: missing required field "LoadBalancerStatus.created_at"`)}
	}
	if _, ok := lbsc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`generated: missing required field "LoadBalancerStatus.updated_at"`)}
	}
	if _, ok := lbsc.mutation.LoadBalancerID(); !ok {
		return &ValidationError{Name: "load_balancer_id", err: errors.New(`generated: missing required field "LoadBalancerStatus.load_balancer_id"`)}
	}
	if _, ok := lbsc.mutation.Source(); !ok {
		return &ValidationError{Name: "source", err: errors.New(`generated: missing required field "LoadBalancerStatus.source"`)}
	}
	if v, ok := lbsc.mutation.Source(); ok {
		if err := loadbalancerstatus.SourceValidator(v); err != nil {
			return &ValidationError{Name: "source", err: fmt.Errorf(`generated: validator failed for field "LoadBalancerStatus.source": %w`, err)}
		}
	}
	if _, ok := lbsc.mutation.LoadBalancerID(); !ok {
		return &ValidationError{Name: "load_balancer", err: errors.New(`generated: missing required edge "LoadBalancerStatus.load_balancer"`)}
	}
	return nil
}

func (lbsc *LoadBalancerStatusCreate) sqlSave(ctx context.Context) (*LoadBalancerStatus, error) {
	if err := lbsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lbsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lbsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*gidx.PrefixedID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	lbsc.mutation.id = &_node.ID
	lbsc.mutation.done = true
	return _node, nil
}

func (lbsc *LoadBalancerStatusCreate) createSpec() (*LoadBalancerStatus, *sqlgraph.CreateSpec) {
	var (
		_node = &LoadBalancerStatus{config: lbsc.config}
		_spec = sqlgraph.NewCreateSpec(loadbalancerstatus.Table, sqlgraph.NewFieldSpec(loadbalancerstatus.FieldID, field.TypeString))
	)
	if id, ok := lbsc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := lbsc.mutation.Namespace(); ok {
		_spec.SetField(loadbalancerstatus.FieldNamespace, field.TypeString, value)
		_node.Namespace = value
	}
	if value, ok := lbsc.mutation.Data(); ok {
		_spec.SetField(loadbalancerstatus.FieldData, field.TypeJSON, value)
		_node.Data = value
	}
	if value, ok := lbsc.mutation.CreatedAt(); ok {
		_spec.SetField(loadbalancerstatus.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := lbsc.mutation.UpdatedAt(); ok {
		_spec.SetField(loadbalancerstatus.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := lbsc.mutation.Source(); ok {
		_spec.SetField(loadbalancerstatus.FieldSource, field.TypeString, value)
		_node.Source = value
	}
	if nodes := lbsc.mutation.LoadBalancerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   loadbalancerstatus.LoadBalancerTable,
			Columns: []string{loadbalancerstatus.LoadBalancerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(loadbalancer.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.LoadBalancerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LoadBalancerStatusCreateBulk is the builder for creating many LoadBalancerStatus entities in bulk.
type LoadBalancerStatusCreateBulk struct {
	config
	builders []*LoadBalancerStatusCreate
}

// Save creates the LoadBalancerStatus entities in the database.
func (lbscb *LoadBalancerStatusCreateBulk) Save(ctx context.Context) ([]*LoadBalancerStatus, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lbscb.builders))
	nodes := make([]*LoadBalancerStatus, len(lbscb.builders))
	mutators := make([]Mutator, len(lbscb.builders))
	for i := range lbscb.builders {
		func(i int, root context.Context) {
			builder := lbscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LoadBalancerStatusMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, lbscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lbscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, lbscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lbscb *LoadBalancerStatusCreateBulk) SaveX(ctx context.Context) []*LoadBalancerStatus {
	v, err := lbscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lbscb *LoadBalancerStatusCreateBulk) Exec(ctx context.Context) error {
	_, err := lbscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lbscb *LoadBalancerStatusCreateBulk) ExecX(ctx context.Context) {
	if err := lbscb.Exec(ctx); err != nil {
		panic(err)
	}
}
