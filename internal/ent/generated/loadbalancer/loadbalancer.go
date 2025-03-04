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

package loadbalancer

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"go.infratographer.com/x/gidx"
)

const (
	// Label holds the string label denoting the loadbalancer type in the database.
	Label = "load_balancer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldTenantID holds the string denoting the tenant_id field in the database.
	FieldTenantID = "tenant_id"
	// FieldLocationID holds the string denoting the location_id field in the database.
	FieldLocationID = "location_id"
	// FieldProviderID holds the string denoting the provider_id field in the database.
	FieldProviderID = "provider_id"
	// EdgeAnnotations holds the string denoting the annotations edge name in mutations.
	EdgeAnnotations = "annotations"
	// EdgeStatuses holds the string denoting the statuses edge name in mutations.
	EdgeStatuses = "statuses"
	// EdgePorts holds the string denoting the ports edge name in mutations.
	EdgePorts = "ports"
	// EdgeProvider holds the string denoting the provider edge name in mutations.
	EdgeProvider = "provider"
	// Table holds the table name of the loadbalancer in the database.
	Table = "load_balancers"
	// AnnotationsTable is the table that holds the annotations relation/edge.
	AnnotationsTable = "load_balancer_annotations"
	// AnnotationsInverseTable is the table name for the LoadBalancerAnnotation entity.
	// It exists in this package in order to avoid circular dependency with the "loadbalancerannotation" package.
	AnnotationsInverseTable = "load_balancer_annotations"
	// AnnotationsColumn is the table column denoting the annotations relation/edge.
	AnnotationsColumn = "load_balancer_id"
	// StatusesTable is the table that holds the statuses relation/edge.
	StatusesTable = "load_balancer_status"
	// StatusesInverseTable is the table name for the LoadBalancerStatus entity.
	// It exists in this package in order to avoid circular dependency with the "loadbalancerstatus" package.
	StatusesInverseTable = "load_balancer_status"
	// StatusesColumn is the table column denoting the statuses relation/edge.
	StatusesColumn = "load_balancer_id"
	// PortsTable is the table that holds the ports relation/edge.
	PortsTable = "ports"
	// PortsInverseTable is the table name for the Port entity.
	// It exists in this package in order to avoid circular dependency with the "port" package.
	PortsInverseTable = "ports"
	// PortsColumn is the table column denoting the ports relation/edge.
	PortsColumn = "load_balancer_id"
	// ProviderTable is the table that holds the provider relation/edge.
	ProviderTable = "load_balancers"
	// ProviderInverseTable is the table name for the Provider entity.
	// It exists in this package in order to avoid circular dependency with the "provider" package.
	ProviderInverseTable = "providers"
	// ProviderColumn is the table column denoting the provider relation/edge.
	ProviderColumn = "provider_id"
)

// Columns holds all SQL columns for loadbalancer fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldTenantID,
	FieldLocationID,
	FieldProviderID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// LocationIDValidator is a validator for the "location_id" field. It is called by the builders before save.
	LocationIDValidator func(string) error
	// ProviderIDValidator is a validator for the "provider_id" field. It is called by the builders before save.
	ProviderIDValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() gidx.PrefixedID
)

// OrderOption defines the ordering options for the LoadBalancer queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByTenantID orders the results by the tenant_id field.
func ByTenantID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTenantID, opts...).ToFunc()
}

// ByLocationID orders the results by the location_id field.
func ByLocationID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocationID, opts...).ToFunc()
}

// ByProviderID orders the results by the provider_id field.
func ByProviderID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProviderID, opts...).ToFunc()
}

// ByAnnotationsCount orders the results by annotations count.
func ByAnnotationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAnnotationsStep(), opts...)
	}
}

// ByAnnotations orders the results by annotations terms.
func ByAnnotations(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAnnotationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByStatusesCount orders the results by statuses count.
func ByStatusesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newStatusesStep(), opts...)
	}
}

// ByStatuses orders the results by statuses terms.
func ByStatuses(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStatusesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPortsCount orders the results by ports count.
func ByPortsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPortsStep(), opts...)
	}
}

// ByPorts orders the results by ports terms.
func ByPorts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPortsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByProviderField orders the results by provider field.
func ByProviderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProviderStep(), sql.OrderByField(field, opts...))
	}
}
func newAnnotationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AnnotationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, AnnotationsTable, AnnotationsColumn),
	)
}
func newStatusesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StatusesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, StatusesTable, StatusesColumn),
	)
}
func newPortsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PortsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, PortsTable, PortsColumn),
	)
}
func newProviderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProviderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ProviderTable, ProviderColumn),
	)
}
