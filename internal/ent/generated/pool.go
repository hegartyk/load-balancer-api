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
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"go.infratographer.com/load-balancer-api/internal/ent/generated/pool"
	"go.infratographer.com/x/gidx"
)

// Pool is the model entity for the Pool schema.
type Pool struct {
	config `json:"-"`
	// ID of the ent.
	ID gidx.PrefixedID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Protocol holds the value of the "protocol" field.
	Protocol pool.Protocol `json:"protocol,omitempty"`
	// TenantID holds the value of the "tenant_id" field.
	TenantID gidx.PrefixedID `json:"tenant_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PoolQuery when eager-loading is set.
	Edges        PoolEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PoolEdges holds the relations/edges for other nodes in the graph.
type PoolEdges struct {
	// Ports holds the value of the ports edge.
	Ports []*Port `json:"ports,omitempty"`
	// Origins holds the value of the origins edge.
	Origins []*Origin `json:"origins,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedPorts   map[string][]*Port
	namedOrigins map[string][]*Origin
}

// PortsOrErr returns the Ports value or an error if the edge
// was not loaded in eager-loading.
func (e PoolEdges) PortsOrErr() ([]*Port, error) {
	if e.loadedTypes[0] {
		return e.Ports, nil
	}
	return nil, &NotLoadedError{edge: "ports"}
}

// OriginsOrErr returns the Origins value or an error if the edge
// was not loaded in eager-loading.
func (e PoolEdges) OriginsOrErr() ([]*Origin, error) {
	if e.loadedTypes[1] {
		return e.Origins, nil
	}
	return nil, &NotLoadedError{edge: "origins"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Pool) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case pool.FieldID, pool.FieldTenantID:
			values[i] = new(gidx.PrefixedID)
		case pool.FieldName, pool.FieldProtocol:
			values[i] = new(sql.NullString)
		case pool.FieldCreatedAt, pool.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Pool fields.
func (po *Pool) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pool.FieldID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				po.ID = *value
			}
		case pool.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				po.CreatedAt = value.Time
			}
		case pool.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				po.UpdatedAt = value.Time
			}
		case pool.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				po.Name = value.String
			}
		case pool.FieldProtocol:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field protocol", values[i])
			} else if value.Valid {
				po.Protocol = pool.Protocol(value.String)
			}
		case pool.FieldTenantID:
			if value, ok := values[i].(*gidx.PrefixedID); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_id", values[i])
			} else if value != nil {
				po.TenantID = *value
			}
		default:
			po.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Pool.
// This includes values selected through modifiers, order, etc.
func (po *Pool) Value(name string) (ent.Value, error) {
	return po.selectValues.Get(name)
}

// QueryPorts queries the "ports" edge of the Pool entity.
func (po *Pool) QueryPorts() *PortQuery {
	return NewPoolClient(po.config).QueryPorts(po)
}

// QueryOrigins queries the "origins" edge of the Pool entity.
func (po *Pool) QueryOrigins() *OriginQuery {
	return NewPoolClient(po.config).QueryOrigins(po)
}

// Update returns a builder for updating this Pool.
// Note that you need to call Pool.Unwrap() before calling this method if this Pool
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *Pool) Update() *PoolUpdateOne {
	return NewPoolClient(po.config).UpdateOne(po)
}

// Unwrap unwraps the Pool entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *Pool) Unwrap() *Pool {
	_tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("generated: Pool is not a transactional entity")
	}
	po.config.driver = _tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *Pool) String() string {
	var builder strings.Builder
	builder.WriteString("Pool(")
	builder.WriteString(fmt.Sprintf("id=%v, ", po.ID))
	builder.WriteString("created_at=")
	builder.WriteString(po.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(po.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(po.Name)
	builder.WriteString(", ")
	builder.WriteString("protocol=")
	builder.WriteString(fmt.Sprintf("%v", po.Protocol))
	builder.WriteString(", ")
	builder.WriteString("tenant_id=")
	builder.WriteString(fmt.Sprintf("%v", po.TenantID))
	builder.WriteByte(')')
	return builder.String()
}

// IsEntity implement fedruntime.Entity
func (po Pool) IsEntity() {}

// NamedPorts returns the Ports named value or an error if the edge was not
// loaded in eager-loading with this name.
func (po *Pool) NamedPorts(name string) ([]*Port, error) {
	if po.Edges.namedPorts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := po.Edges.namedPorts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (po *Pool) appendNamedPorts(name string, edges ...*Port) {
	if po.Edges.namedPorts == nil {
		po.Edges.namedPorts = make(map[string][]*Port)
	}
	if len(edges) == 0 {
		po.Edges.namedPorts[name] = []*Port{}
	} else {
		po.Edges.namedPorts[name] = append(po.Edges.namedPorts[name], edges...)
	}
}

// NamedOrigins returns the Origins named value or an error if the edge was not
// loaded in eager-loading with this name.
func (po *Pool) NamedOrigins(name string) ([]*Origin, error) {
	if po.Edges.namedOrigins == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := po.Edges.namedOrigins[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (po *Pool) appendNamedOrigins(name string, edges ...*Origin) {
	if po.Edges.namedOrigins == nil {
		po.Edges.namedOrigins = make(map[string][]*Origin)
	}
	if len(edges) == 0 {
		po.Edges.namedOrigins[name] = []*Origin{}
	} else {
		po.Edges.namedOrigins[name] = append(po.Edges.namedOrigins[name], edges...)
	}
}

// Pools is a parsable slice of Pool.
type Pools []*Pool
