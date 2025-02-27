// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"pinares/ent/apartment"
	"pinares/ent/predicate"
	"pinares/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ApartmentUpdate is the builder for updating Apartment entities.
type ApartmentUpdate struct {
	config
	hooks    []Hook
	mutation *ApartmentMutation
}

// Where appends a list predicates to the ApartmentUpdate builder.
func (au *ApartmentUpdate) Where(ps ...predicate.Apartment) *ApartmentUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetArea sets the "area" field.
func (au *ApartmentUpdate) SetArea(i int16) *ApartmentUpdate {
	au.mutation.ResetArea()
	au.mutation.SetArea(i)
	return au
}

// SetNillableArea sets the "area" field if the given value is not nil.
func (au *ApartmentUpdate) SetNillableArea(i *int16) *ApartmentUpdate {
	if i != nil {
		au.SetArea(*i)
	}
	return au
}

// AddArea adds i to the "area" field.
func (au *ApartmentUpdate) AddArea(i int16) *ApartmentUpdate {
	au.mutation.AddArea(i)
	return au
}

// SetFloor sets the "floor" field.
func (au *ApartmentUpdate) SetFloor(i int16) *ApartmentUpdate {
	au.mutation.ResetFloor()
	au.mutation.SetFloor(i)
	return au
}

// SetNillableFloor sets the "floor" field if the given value is not nil.
func (au *ApartmentUpdate) SetNillableFloor(i *int16) *ApartmentUpdate {
	if i != nil {
		au.SetFloor(*i)
	}
	return au
}

// AddFloor adds i to the "floor" field.
func (au *ApartmentUpdate) AddFloor(i int16) *ApartmentUpdate {
	au.mutation.AddFloor(i)
	return au
}

// SetNumber sets the "number" field.
func (au *ApartmentUpdate) SetNumber(i int16) *ApartmentUpdate {
	au.mutation.ResetNumber()
	au.mutation.SetNumber(i)
	return au
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (au *ApartmentUpdate) SetNillableNumber(i *int16) *ApartmentUpdate {
	if i != nil {
		au.SetNumber(*i)
	}
	return au
}

// AddNumber adds i to the "number" field.
func (au *ApartmentUpdate) AddNumber(i int16) *ApartmentUpdate {
	au.mutation.AddNumber(i)
	return au
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (au *ApartmentUpdate) SetOwnerID(id int) *ApartmentUpdate {
	au.mutation.SetOwnerID(id)
	return au
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (au *ApartmentUpdate) SetNillableOwnerID(id *int) *ApartmentUpdate {
	if id != nil {
		au = au.SetOwnerID(*id)
	}
	return au
}

// SetOwner sets the "owner" edge to the User entity.
func (au *ApartmentUpdate) SetOwner(u *User) *ApartmentUpdate {
	return au.SetOwnerID(u.ID)
}

// Mutation returns the ApartmentMutation object of the builder.
func (au *ApartmentUpdate) Mutation() *ApartmentMutation {
	return au.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (au *ApartmentUpdate) ClearOwner() *ApartmentUpdate {
	au.mutation.ClearOwner()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ApartmentUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *ApartmentUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ApartmentUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ApartmentUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *ApartmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(apartment.Table, apartment.Columns, sqlgraph.NewFieldSpec(apartment.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Area(); ok {
		_spec.SetField(apartment.FieldArea, field.TypeInt16, value)
	}
	if value, ok := au.mutation.AddedArea(); ok {
		_spec.AddField(apartment.FieldArea, field.TypeInt16, value)
	}
	if value, ok := au.mutation.Floor(); ok {
		_spec.SetField(apartment.FieldFloor, field.TypeInt16, value)
	}
	if value, ok := au.mutation.AddedFloor(); ok {
		_spec.AddField(apartment.FieldFloor, field.TypeInt16, value)
	}
	if value, ok := au.mutation.Number(); ok {
		_spec.SetField(apartment.FieldNumber, field.TypeInt16, value)
	}
	if value, ok := au.mutation.AddedNumber(); ok {
		_spec.AddField(apartment.FieldNumber, field.TypeInt16, value)
	}
	if au.mutation.OwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apartment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// ApartmentUpdateOne is the builder for updating a single Apartment entity.
type ApartmentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ApartmentMutation
}

// SetArea sets the "area" field.
func (auo *ApartmentUpdateOne) SetArea(i int16) *ApartmentUpdateOne {
	auo.mutation.ResetArea()
	auo.mutation.SetArea(i)
	return auo
}

// SetNillableArea sets the "area" field if the given value is not nil.
func (auo *ApartmentUpdateOne) SetNillableArea(i *int16) *ApartmentUpdateOne {
	if i != nil {
		auo.SetArea(*i)
	}
	return auo
}

// AddArea adds i to the "area" field.
func (auo *ApartmentUpdateOne) AddArea(i int16) *ApartmentUpdateOne {
	auo.mutation.AddArea(i)
	return auo
}

// SetFloor sets the "floor" field.
func (auo *ApartmentUpdateOne) SetFloor(i int16) *ApartmentUpdateOne {
	auo.mutation.ResetFloor()
	auo.mutation.SetFloor(i)
	return auo
}

// SetNillableFloor sets the "floor" field if the given value is not nil.
func (auo *ApartmentUpdateOne) SetNillableFloor(i *int16) *ApartmentUpdateOne {
	if i != nil {
		auo.SetFloor(*i)
	}
	return auo
}

// AddFloor adds i to the "floor" field.
func (auo *ApartmentUpdateOne) AddFloor(i int16) *ApartmentUpdateOne {
	auo.mutation.AddFloor(i)
	return auo
}

// SetNumber sets the "number" field.
func (auo *ApartmentUpdateOne) SetNumber(i int16) *ApartmentUpdateOne {
	auo.mutation.ResetNumber()
	auo.mutation.SetNumber(i)
	return auo
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (auo *ApartmentUpdateOne) SetNillableNumber(i *int16) *ApartmentUpdateOne {
	if i != nil {
		auo.SetNumber(*i)
	}
	return auo
}

// AddNumber adds i to the "number" field.
func (auo *ApartmentUpdateOne) AddNumber(i int16) *ApartmentUpdateOne {
	auo.mutation.AddNumber(i)
	return auo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (auo *ApartmentUpdateOne) SetOwnerID(id int) *ApartmentUpdateOne {
	auo.mutation.SetOwnerID(id)
	return auo
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (auo *ApartmentUpdateOne) SetNillableOwnerID(id *int) *ApartmentUpdateOne {
	if id != nil {
		auo = auo.SetOwnerID(*id)
	}
	return auo
}

// SetOwner sets the "owner" edge to the User entity.
func (auo *ApartmentUpdateOne) SetOwner(u *User) *ApartmentUpdateOne {
	return auo.SetOwnerID(u.ID)
}

// Mutation returns the ApartmentMutation object of the builder.
func (auo *ApartmentUpdateOne) Mutation() *ApartmentMutation {
	return auo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (auo *ApartmentUpdateOne) ClearOwner() *ApartmentUpdateOne {
	auo.mutation.ClearOwner()
	return auo
}

// Where appends a list predicates to the ApartmentUpdate builder.
func (auo *ApartmentUpdateOne) Where(ps ...predicate.Apartment) *ApartmentUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *ApartmentUpdateOne) Select(field string, fields ...string) *ApartmentUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Apartment entity.
func (auo *ApartmentUpdateOne) Save(ctx context.Context) (*Apartment, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ApartmentUpdateOne) SaveX(ctx context.Context) *Apartment {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ApartmentUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ApartmentUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *ApartmentUpdateOne) sqlSave(ctx context.Context) (_node *Apartment, err error) {
	_spec := sqlgraph.NewUpdateSpec(apartment.Table, apartment.Columns, sqlgraph.NewFieldSpec(apartment.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Apartment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, apartment.FieldID)
		for _, f := range fields {
			if !apartment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != apartment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Area(); ok {
		_spec.SetField(apartment.FieldArea, field.TypeInt16, value)
	}
	if value, ok := auo.mutation.AddedArea(); ok {
		_spec.AddField(apartment.FieldArea, field.TypeInt16, value)
	}
	if value, ok := auo.mutation.Floor(); ok {
		_spec.SetField(apartment.FieldFloor, field.TypeInt16, value)
	}
	if value, ok := auo.mutation.AddedFloor(); ok {
		_spec.AddField(apartment.FieldFloor, field.TypeInt16, value)
	}
	if value, ok := auo.mutation.Number(); ok {
		_spec.SetField(apartment.FieldNumber, field.TypeInt16, value)
	}
	if value, ok := auo.mutation.AddedNumber(); ok {
		_spec.AddField(apartment.FieldNumber, field.TypeInt16, value)
	}
	if auo.mutation.OwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Apartment{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apartment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
