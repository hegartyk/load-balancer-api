// Code generated by SQLBoiler 4.14.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// LoadBalancer is an object representing the database table.
type LoadBalancer struct {
	CreatedAt        time.Time `query:"created_at" param:"created_at" boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt        time.Time `query:"updated_at" param:"updated_at" boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	DeletedAt        null.Time `query:"deleted_at" param:"deleted_at" boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	LoadBalancerID   string    `query:"load_balancer_id" param:"load_balancer_id" boil:"load_balancer_id" json:"load_balancer_id" toml:"load_balancer_id" yaml:"load_balancer_id"`
	LocationID       string    `query:"location_id" param:"location_id" boil:"location_id" json:"location_id" toml:"location_id" yaml:"location_id"`
	TenantID         string    `query:"tenant_id" param:"tenant_id" boil:"tenant_id" json:"tenant_id" toml:"tenant_id" yaml:"tenant_id"`
	Name             string    `query:"name" param:"name" boil:"name" json:"name" toml:"name" yaml:"name"`
	Slug             string    `query:"slug" param:"slug" boil:"slug" json:"slug" toml:"slug" yaml:"slug"`
	LoadBalancerSize string    `query:"load_balancer_size" param:"load_balancer_size" boil:"load_balancer_size" json:"load_balancer_size" toml:"load_balancer_size" yaml:"load_balancer_size"`
	LoadBalancerType string    `query:"load_balancer_type" param:"load_balancer_type" boil:"load_balancer_type" json:"load_balancer_type" toml:"load_balancer_type" yaml:"load_balancer_type"`
	StateChangedAt   null.Time `query:"state_changed_at" param:"state_changed_at" boil:"state_changed_at" json:"state_changed_at,omitempty" toml:"state_changed_at" yaml:"state_changed_at,omitempty"`
	CurrentState     string    `query:"current_state" param:"current_state" boil:"current_state" json:"current_state" toml:"current_state" yaml:"current_state"`
	PreviousState    string    `query:"previous_state" param:"previous_state" boil:"previous_state" json:"previous_state" toml:"previous_state" yaml:"previous_state"`
	IPAddressID      string    `query:"ip_address_id" param:"ip_address_id" boil:"ip_address_id" json:"ip_address_id" toml:"ip_address_id" yaml:"ip_address_id"`

	R *loadBalancerR `query:"-" param:"-" boil:"-" json:"-" toml:"-" yaml:"-"`
	L loadBalancerL  `query:"-" param:"-" boil:"-" json:"-" toml:"-" yaml:"-"`
}

var LoadBalancerColumns = struct {
	CreatedAt        string
	UpdatedAt        string
	DeletedAt        string
	LoadBalancerID   string
	LocationID       string
	TenantID         string
	Name             string
	Slug             string
	LoadBalancerSize string
	LoadBalancerType string
	StateChangedAt   string
	CurrentState     string
	PreviousState    string
	IPAddressID      string
}{
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	DeletedAt:        "deleted_at",
	LoadBalancerID:   "load_balancer_id",
	LocationID:       "location_id",
	TenantID:         "tenant_id",
	Name:             "name",
	Slug:             "slug",
	LoadBalancerSize: "load_balancer_size",
	LoadBalancerType: "load_balancer_type",
	StateChangedAt:   "state_changed_at",
	CurrentState:     "current_state",
	PreviousState:    "previous_state",
	IPAddressID:      "ip_address_id",
}

var LoadBalancerTableColumns = struct {
	CreatedAt        string
	UpdatedAt        string
	DeletedAt        string
	LoadBalancerID   string
	LocationID       string
	TenantID         string
	Name             string
	Slug             string
	LoadBalancerSize string
	LoadBalancerType string
	StateChangedAt   string
	CurrentState     string
	PreviousState    string
	IPAddressID      string
}{
	CreatedAt:        "load_balancers.created_at",
	UpdatedAt:        "load_balancers.updated_at",
	DeletedAt:        "load_balancers.deleted_at",
	LoadBalancerID:   "load_balancers.load_balancer_id",
	LocationID:       "load_balancers.location_id",
	TenantID:         "load_balancers.tenant_id",
	Name:             "load_balancers.name",
	Slug:             "load_balancers.slug",
	LoadBalancerSize: "load_balancers.load_balancer_size",
	LoadBalancerType: "load_balancers.load_balancer_type",
	StateChangedAt:   "load_balancers.state_changed_at",
	CurrentState:     "load_balancers.current_state",
	PreviousState:    "load_balancers.previous_state",
	IPAddressID:      "load_balancers.ip_address_id",
}

// Generated where

var LoadBalancerWhere = struct {
	CreatedAt        whereHelpertime_Time
	UpdatedAt        whereHelpertime_Time
	DeletedAt        whereHelpernull_Time
	LoadBalancerID   whereHelperstring
	LocationID       whereHelperstring
	TenantID         whereHelperstring
	Name             whereHelperstring
	Slug             whereHelperstring
	LoadBalancerSize whereHelperstring
	LoadBalancerType whereHelperstring
	StateChangedAt   whereHelpernull_Time
	CurrentState     whereHelperstring
	PreviousState    whereHelperstring
	IPAddressID      whereHelperstring
}{
	CreatedAt:        whereHelpertime_Time{field: "\"load_balancers\".\"created_at\""},
	UpdatedAt:        whereHelpertime_Time{field: "\"load_balancers\".\"updated_at\""},
	DeletedAt:        whereHelpernull_Time{field: "\"load_balancers\".\"deleted_at\""},
	LoadBalancerID:   whereHelperstring{field: "\"load_balancers\".\"load_balancer_id\""},
	LocationID:       whereHelperstring{field: "\"load_balancers\".\"location_id\""},
	TenantID:         whereHelperstring{field: "\"load_balancers\".\"tenant_id\""},
	Name:             whereHelperstring{field: "\"load_balancers\".\"name\""},
	Slug:             whereHelperstring{field: "\"load_balancers\".\"slug\""},
	LoadBalancerSize: whereHelperstring{field: "\"load_balancers\".\"load_balancer_size\""},
	LoadBalancerType: whereHelperstring{field: "\"load_balancers\".\"load_balancer_type\""},
	StateChangedAt:   whereHelpernull_Time{field: "\"load_balancers\".\"state_changed_at\""},
	CurrentState:     whereHelperstring{field: "\"load_balancers\".\"current_state\""},
	PreviousState:    whereHelperstring{field: "\"load_balancers\".\"previous_state\""},
	IPAddressID:      whereHelperstring{field: "\"load_balancers\".\"ip_address_id\""},
}

// LoadBalancerRels is where relationship names are stored.
var LoadBalancerRels = struct {
	Frontends string
}{
	Frontends: "Frontends",
}

// loadBalancerR is where relationships are stored.
type loadBalancerR struct {
	Frontends FrontendSlice `query:"Frontends" param:"Frontends" boil:"Frontends" json:"Frontends" toml:"Frontends" yaml:"Frontends"`
}

// NewStruct creates a new relationship struct
func (*loadBalancerR) NewStruct() *loadBalancerR {
	return &loadBalancerR{}
}

func (r *loadBalancerR) GetFrontends() FrontendSlice {
	if r == nil {
		return nil
	}
	return r.Frontends
}

// loadBalancerL is where Load methods for each relationship are stored.
type loadBalancerL struct{}

var (
	loadBalancerAllColumns            = []string{"created_at", "updated_at", "deleted_at", "load_balancer_id", "location_id", "tenant_id", "name", "slug", "load_balancer_size", "load_balancer_type", "state_changed_at", "current_state", "previous_state", "ip_address_id"}
	loadBalancerColumnsWithoutDefault = []string{"location_id", "tenant_id", "name", "slug", "load_balancer_size", "load_balancer_type", "current_state", "previous_state", "ip_address_id"}
	loadBalancerColumnsWithDefault    = []string{"created_at", "updated_at", "deleted_at", "load_balancer_id", "state_changed_at"}
	loadBalancerPrimaryKeyColumns     = []string{"load_balancer_id"}
	loadBalancerGeneratedColumns      = []string{}
)

type (
	// LoadBalancerSlice is an alias for a slice of pointers to LoadBalancer.
	// This should almost always be used instead of []LoadBalancer.
	LoadBalancerSlice []*LoadBalancer
	// LoadBalancerHook is the signature for custom LoadBalancer hook methods
	LoadBalancerHook func(context.Context, boil.ContextExecutor, *LoadBalancer) error

	loadBalancerQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	loadBalancerType                 = reflect.TypeOf(&LoadBalancer{})
	loadBalancerMapping              = queries.MakeStructMapping(loadBalancerType)
	loadBalancerPrimaryKeyMapping, _ = queries.BindMapping(loadBalancerType, loadBalancerMapping, loadBalancerPrimaryKeyColumns)
	loadBalancerInsertCacheMut       sync.RWMutex
	loadBalancerInsertCache          = make(map[string]insertCache)
	loadBalancerUpdateCacheMut       sync.RWMutex
	loadBalancerUpdateCache          = make(map[string]updateCache)
	loadBalancerUpsertCacheMut       sync.RWMutex
	loadBalancerUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var loadBalancerAfterSelectHooks []LoadBalancerHook

var loadBalancerBeforeInsertHooks []LoadBalancerHook
var loadBalancerAfterInsertHooks []LoadBalancerHook

var loadBalancerBeforeUpdateHooks []LoadBalancerHook
var loadBalancerAfterUpdateHooks []LoadBalancerHook

var loadBalancerBeforeDeleteHooks []LoadBalancerHook
var loadBalancerAfterDeleteHooks []LoadBalancerHook

var loadBalancerBeforeUpsertHooks []LoadBalancerHook
var loadBalancerAfterUpsertHooks []LoadBalancerHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *LoadBalancer) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loadBalancerAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *LoadBalancer) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loadBalancerBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *LoadBalancer) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loadBalancerAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *LoadBalancer) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loadBalancerBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *LoadBalancer) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loadBalancerAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *LoadBalancer) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loadBalancerBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *LoadBalancer) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loadBalancerAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *LoadBalancer) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loadBalancerBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *LoadBalancer) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range loadBalancerAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddLoadBalancerHook registers your hook function for all future operations.
func AddLoadBalancerHook(hookPoint boil.HookPoint, loadBalancerHook LoadBalancerHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		loadBalancerAfterSelectHooks = append(loadBalancerAfterSelectHooks, loadBalancerHook)
	case boil.BeforeInsertHook:
		loadBalancerBeforeInsertHooks = append(loadBalancerBeforeInsertHooks, loadBalancerHook)
	case boil.AfterInsertHook:
		loadBalancerAfterInsertHooks = append(loadBalancerAfterInsertHooks, loadBalancerHook)
	case boil.BeforeUpdateHook:
		loadBalancerBeforeUpdateHooks = append(loadBalancerBeforeUpdateHooks, loadBalancerHook)
	case boil.AfterUpdateHook:
		loadBalancerAfterUpdateHooks = append(loadBalancerAfterUpdateHooks, loadBalancerHook)
	case boil.BeforeDeleteHook:
		loadBalancerBeforeDeleteHooks = append(loadBalancerBeforeDeleteHooks, loadBalancerHook)
	case boil.AfterDeleteHook:
		loadBalancerAfterDeleteHooks = append(loadBalancerAfterDeleteHooks, loadBalancerHook)
	case boil.BeforeUpsertHook:
		loadBalancerBeforeUpsertHooks = append(loadBalancerBeforeUpsertHooks, loadBalancerHook)
	case boil.AfterUpsertHook:
		loadBalancerAfterUpsertHooks = append(loadBalancerAfterUpsertHooks, loadBalancerHook)
	}
}

// One returns a single loadBalancer record from the query.
func (q loadBalancerQuery) One(ctx context.Context, exec boil.ContextExecutor) (*LoadBalancer, error) {
	o := &LoadBalancer{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to execute a one query for load_balancers")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all LoadBalancer records from the query.
func (q loadBalancerQuery) All(ctx context.Context, exec boil.ContextExecutor) (LoadBalancerSlice, error) {
	var o []*LoadBalancer

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to LoadBalancer slice")
	}

	if len(loadBalancerAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all LoadBalancer records in the query.
func (q loadBalancerQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count load_balancers rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q loadBalancerQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if load_balancers exists")
	}

	return count > 0, nil
}

// Frontends retrieves all the frontend's Frontends with an executor.
func (o *LoadBalancer) Frontends(mods ...qm.QueryMod) frontendQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"frontends\".\"load_balancer_id\"=?", o.LoadBalancerID),
	)

	return Frontends(queryMods...)
}

// LoadFrontends allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (loadBalancerL) LoadFrontends(ctx context.Context, e boil.ContextExecutor, singular bool, maybeLoadBalancer interface{}, mods queries.Applicator) error {
	var slice []*LoadBalancer
	var object *LoadBalancer

	if singular {
		var ok bool
		object, ok = maybeLoadBalancer.(*LoadBalancer)
		if !ok {
			object = new(LoadBalancer)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeLoadBalancer)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeLoadBalancer))
			}
		}
	} else {
		s, ok := maybeLoadBalancer.(*[]*LoadBalancer)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeLoadBalancer)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeLoadBalancer))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &loadBalancerR{}
		}
		args = append(args, object.LoadBalancerID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &loadBalancerR{}
			}

			for _, a := range args {
				if a == obj.LoadBalancerID {
					continue Outer
				}
			}

			args = append(args, obj.LoadBalancerID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`frontends`),
		qm.WhereIn(`frontends.load_balancer_id in ?`, args...),
		qmhelper.WhereIsNull(`frontends.deleted_at`),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load frontends")
	}

	var resultSlice []*Frontend
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice frontends")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on frontends")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for frontends")
	}

	if len(frontendAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Frontends = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &frontendR{}
			}
			foreign.R.LoadBalancer = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.LoadBalancerID == foreign.LoadBalancerID {
				local.R.Frontends = append(local.R.Frontends, foreign)
				if foreign.R == nil {
					foreign.R = &frontendR{}
				}
				foreign.R.LoadBalancer = local
				break
			}
		}
	}

	return nil
}

// AddFrontends adds the given related objects to the existing relationships
// of the load_balancer, optionally inserting them as new records.
// Appends related to o.R.Frontends.
// Sets related.R.LoadBalancer appropriately.
func (o *LoadBalancer) AddFrontends(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Frontend) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.LoadBalancerID = o.LoadBalancerID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"frontends\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"load_balancer_id"}),
				strmangle.WhereClause("\"", "\"", 2, frontendPrimaryKeyColumns),
			)
			values := []interface{}{o.LoadBalancerID, rel.FrontendID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.LoadBalancerID = o.LoadBalancerID
		}
	}

	if o.R == nil {
		o.R = &loadBalancerR{
			Frontends: related,
		}
	} else {
		o.R.Frontends = append(o.R.Frontends, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &frontendR{
				LoadBalancer: o,
			}
		} else {
			rel.R.LoadBalancer = o
		}
	}
	return nil
}

// LoadBalancers retrieves all the records using an executor.
func LoadBalancers(mods ...qm.QueryMod) loadBalancerQuery {
	mods = append(mods, qm.From("\"load_balancers\""), qmhelper.WhereIsNull("\"load_balancers\".\"deleted_at\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"load_balancers\".*"})
	}

	return loadBalancerQuery{q}
}

// FindLoadBalancer retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindLoadBalancer(ctx context.Context, exec boil.ContextExecutor, loadBalancerID string, selectCols ...string) (*LoadBalancer, error) {
	loadBalancerObj := &LoadBalancer{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"load_balancers\" where \"load_balancer_id\"=$1 and \"deleted_at\" is null", sel,
	)

	q := queries.Raw(query, loadBalancerID)

	err := q.Bind(ctx, exec, loadBalancerObj)
	if err != nil {
		return nil, errors.Wrap(err, "models: unable to select from load_balancers")
	}

	if err = loadBalancerObj.doAfterSelectHooks(ctx, exec); err != nil {
		return loadBalancerObj, err
	}

	return loadBalancerObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *LoadBalancer) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no load_balancers provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(loadBalancerColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	loadBalancerInsertCacheMut.RLock()
	cache, cached := loadBalancerInsertCache[key]
	loadBalancerInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			loadBalancerAllColumns,
			loadBalancerColumnsWithDefault,
			loadBalancerColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(loadBalancerType, loadBalancerMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(loadBalancerType, loadBalancerMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"load_balancers\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"load_balancers\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into load_balancers")
	}

	if !cached {
		loadBalancerInsertCacheMut.Lock()
		loadBalancerInsertCache[key] = cache
		loadBalancerInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the LoadBalancer.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *LoadBalancer) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	loadBalancerUpdateCacheMut.RLock()
	cache, cached := loadBalancerUpdateCache[key]
	loadBalancerUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			loadBalancerAllColumns,
			loadBalancerPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update load_balancers, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"load_balancers\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, loadBalancerPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(loadBalancerType, loadBalancerMapping, append(wl, loadBalancerPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update load_balancers row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for load_balancers")
	}

	if !cached {
		loadBalancerUpdateCacheMut.Lock()
		loadBalancerUpdateCache[key] = cache
		loadBalancerUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q loadBalancerQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for load_balancers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for load_balancers")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o LoadBalancerSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), loadBalancerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"load_balancers\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, loadBalancerPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in loadBalancer slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all loadBalancer")
	}
	return rowsAff, nil
}

// Delete deletes a single LoadBalancer record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *LoadBalancer) Delete(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no LoadBalancer provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	var (
		sql  string
		args []interface{}
	)
	if hardDelete {
		args = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), loadBalancerPrimaryKeyMapping)
		sql = "DELETE FROM \"load_balancers\" WHERE \"load_balancer_id\"=$1"
	} else {
		currTime := time.Now().In(boil.GetLocation())
		o.DeletedAt = null.TimeFrom(currTime)
		wl := []string{"deleted_at"}
		sql = fmt.Sprintf("UPDATE \"load_balancers\" SET %s WHERE \"load_balancer_id\"=$2",
			strmangle.SetParamNames("\"", "\"", 1, wl),
		)
		valueMapping, err := queries.BindMapping(loadBalancerType, loadBalancerMapping, append(wl, loadBalancerPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
		args = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), valueMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from load_balancers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for load_balancers")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q loadBalancerQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no loadBalancerQuery provided for delete all")
	}

	if hardDelete {
		queries.SetDelete(q.Query)
	} else {
		currTime := time.Now().In(boil.GetLocation())
		queries.SetUpdate(q.Query, M{"deleted_at": currTime})
	}

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from load_balancers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for load_balancers")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o LoadBalancerSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(loadBalancerBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var (
		sql  string
		args []interface{}
	)
	if hardDelete {
		for _, obj := range o {
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), loadBalancerPrimaryKeyMapping)
			args = append(args, pkeyArgs...)
		}
		sql = "DELETE FROM \"load_balancers\" WHERE " +
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, loadBalancerPrimaryKeyColumns, len(o))
	} else {
		currTime := time.Now().In(boil.GetLocation())
		for _, obj := range o {
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), loadBalancerPrimaryKeyMapping)
			args = append(args, pkeyArgs...)
			obj.DeletedAt = null.TimeFrom(currTime)
		}
		wl := []string{"deleted_at"}
		sql = fmt.Sprintf("UPDATE \"load_balancers\" SET %s WHERE "+
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 2, loadBalancerPrimaryKeyColumns, len(o)),
			strmangle.SetParamNames("\"", "\"", 1, wl),
		)
		args = append([]interface{}{currTime}, args...)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from loadBalancer slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for load_balancers")
	}

	if len(loadBalancerAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *LoadBalancer) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindLoadBalancer(ctx, exec, o.LoadBalancerID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LoadBalancerSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := LoadBalancerSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), loadBalancerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"load_balancers\".* FROM \"load_balancers\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, loadBalancerPrimaryKeyColumns, len(*o)) +
		"and \"deleted_at\" is null"

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in LoadBalancerSlice")
	}

	*o = slice

	return nil
}

// LoadBalancerExists checks if the LoadBalancer row exists.
func LoadBalancerExists(ctx context.Context, exec boil.ContextExecutor, loadBalancerID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"load_balancers\" where \"load_balancer_id\"=$1 and \"deleted_at\" is null limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, loadBalancerID)
	}
	row := exec.QueryRowContext(ctx, sql, loadBalancerID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if load_balancers exists")
	}

	return exists, nil
}

// Exists checks if the LoadBalancer row exists.
func (o *LoadBalancer) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return LoadBalancerExists(ctx, exec, o.LoadBalancerID)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *LoadBalancer) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no load_balancers provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(loadBalancerColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	loadBalancerUpsertCacheMut.RLock()
	cache, cached := loadBalancerUpsertCache[key]
	loadBalancerUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			loadBalancerAllColumns,
			loadBalancerColumnsWithDefault,
			loadBalancerColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			loadBalancerAllColumns,
			loadBalancerPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert load_balancers, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(loadBalancerPrimaryKeyColumns))
			copy(conflict, loadBalancerPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryCockroachDB(dialect, "\"load_balancers\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(loadBalancerType, loadBalancerMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(loadBalancerType, loadBalancerMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		_, _ = fmt.Fprintln(boil.DebugWriter, cache.query)
		_, _ = fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // CockcorachDB doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert load_balancers")
	}

	if !cached {
		loadBalancerUpsertCacheMut.Lock()
		loadBalancerUpsertCache[key] = cache
		loadBalancerUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}
