// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"pinares/ent/apartment"
	"pinares/ent/user"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ApartmentCreate is the builder for creating a Apartment entity.
type ApartmentCreate struct {
	config
	mutation *ApartmentMutation
	hooks    []Hook
}

// SetArea sets the "area" field.
func (ac *ApartmentCreate) SetArea(i int16) *ApartmentCreate {
	ac.mutation.SetArea(i)
	return ac
}

// SetFloor sets the "floor" field.
func (ac *ApartmentCreate) SetFloor(i int16) *ApartmentCreate {
	ac.mutation.SetFloor(i)
	return ac
}

// SetNumber sets the "number" field.
func (ac *ApartmentCreate) SetNumber(i int16) *ApartmentCreate {
	ac.mutation.SetNumber(i)
	return ac
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (ac *ApartmentCreate) SetOwnerID(id int) *ApartmentCreate {
	ac.mutation.SetOwnerID(id)
	return ac
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (ac *ApartmentCreate) SetNillableOwnerID(id *int) *ApartmentCreate {
	if id != nil {
		ac = ac.SetOwnerID(*id)
	}
	return ac
}

// SetOwner sets the "owner" edge to the User entity.
func (ac *ApartmentCreate) SetOwner(u *User) *ApartmentCreate {
	return ac.SetOwnerID(u.ID)
}

// Mutation returns the ApartmentMutation object of the builder.
func (ac *ApartmentCreate) Mutation() *ApartmentMutation {
	return ac.mutation
}

// Save creates the Apartment in the database.
func (ac *ApartmentCreate) Save(ctx context.Context) (*Apartment, error) {
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ApartmentCreate) SaveX(ctx context.Context) *Apartment {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *ApartmentCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *ApartmentCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *ApartmentCreate) check() error {
	if _, ok := ac.mutation.Area(); !ok {
		return &ValidationError{Name: "area", err: errors.New(`ent: missing required field "Apartment.area"`)}
	}
	if _, ok := ac.mutation.Floor(); !ok {
		return &ValidationError{Name: "floor", err: errors.New(`ent: missing required field "Apartment.floor"`)}
	}
	if _, ok := ac.mutation.Number(); !ok {
		return &ValidationError{Name: "number", err: errors.New(`ent: missing required field "Apartment.number"`)}
	}
	return nil
}

func (ac *ApartmentCreate) sqlSave(ctx context.Context) (*Apartment, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *ApartmentCreate) createSpec() (*Apartment, *sqlgraph.CreateSpec) {
	var (
		_node = &Apartment{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(apartment.Table, sqlgraph.NewFieldSpec(apartment.FieldID, field.TypeInt))
	)
	if value, ok := ac.mutation.Area(); ok {
		_spec.SetField(apartment.FieldArea, field.TypeInt16, value)
		_node.Area = value
	}
	if value, ok := ac.mutation.Floor(); ok {
		_spec.SetField(apartment.FieldFloor, field.TypeInt16, value)
		_node.Floor = value
	}
	if value, ok := ac.mutation.Number(); ok {
		_spec.SetField(apartment.FieldNumber, field.TypeInt16, value)
		_node.Number = value
	}
	if nodes := ac.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   apartment.OwnerTable,
			Columns: []string{apartment.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.apartment_owner = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ApartmentCreateBulk is the builder for creating many Apartment entities in bulk.
type ApartmentCreateBulk struct {
	config
	err      error
	builders []*ApartmentCreate
}

// Save creates the Apartment entities in the database.
func (acb *ApartmentCreateBulk) Save(ctx context.Context) ([]*Apartment, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Apartment, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ApartmentMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *ApartmentCreateBulk) SaveX(ctx context.Context) []*Apartment {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *ApartmentCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *ApartmentCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
