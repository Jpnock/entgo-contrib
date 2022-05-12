// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/contrib/entproto/internal/entprototest/ent/messagewithid"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageWithIDCreate is the builder for creating a MessageWithID entity.
type MessageWithIDCreate struct {
	config
	mutation *MessageWithIDMutation
	hooks    []Hook
}

// SetID sets the "id" field.
func (mwic *MessageWithIDCreate) SetID(i int32) *MessageWithIDCreate {
	mwic.mutation.SetID(i)
	return mwic
}

// Mutation returns the MessageWithIDMutation object of the builder.
func (mwic *MessageWithIDCreate) Mutation() *MessageWithIDMutation {
	return mwic.mutation
}

// Save creates the MessageWithID in the database.
func (mwic *MessageWithIDCreate) Save(ctx context.Context) (*MessageWithID, error) {
	var (
		err  error
		node *MessageWithID
	)
	if len(mwic.hooks) == 0 {
		if err = mwic.check(); err != nil {
			return nil, err
		}
		node, err = mwic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MessageWithIDMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mwic.check(); err != nil {
				return nil, err
			}
			mwic.mutation = mutation
			if node, err = mwic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mwic.hooks) - 1; i >= 0; i-- {
			if mwic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mwic.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, mwic.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*MessageWithID)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MessageWithIDMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mwic *MessageWithIDCreate) SaveX(ctx context.Context) *MessageWithID {
	v, err := mwic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mwic *MessageWithIDCreate) Exec(ctx context.Context) error {
	_, err := mwic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mwic *MessageWithIDCreate) ExecX(ctx context.Context) {
	if err := mwic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mwic *MessageWithIDCreate) check() error {
	return nil
}

func (mwic *MessageWithIDCreate) sqlSave(ctx context.Context) (*MessageWithID, error) {
	_node, _spec := mwic.createSpec()
	if err := sqlgraph.CreateNode(ctx, mwic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	return _node, nil
}

func (mwic *MessageWithIDCreate) createSpec() (*MessageWithID, *sqlgraph.CreateSpec) {
	var (
		_node = &MessageWithID{config: mwic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: messagewithid.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: messagewithid.FieldID,
			},
		}
	)
	if id, ok := mwic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	return _node, _spec
}

// MessageWithIDCreateBulk is the builder for creating many MessageWithID entities in bulk.
type MessageWithIDCreateBulk struct {
	config
	builders []*MessageWithIDCreate
}

// Save creates the MessageWithID entities in the database.
func (mwicb *MessageWithIDCreateBulk) Save(ctx context.Context) ([]*MessageWithID, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mwicb.builders))
	nodes := make([]*MessageWithID, len(mwicb.builders))
	mutators := make([]Mutator, len(mwicb.builders))
	for i := range mwicb.builders {
		func(i int, root context.Context) {
			builder := mwicb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MessageWithIDMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mwicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mwicb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int32(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mwicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mwicb *MessageWithIDCreateBulk) SaveX(ctx context.Context) []*MessageWithID {
	v, err := mwicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mwicb *MessageWithIDCreateBulk) Exec(ctx context.Context) error {
	_, err := mwicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mwicb *MessageWithIDCreateBulk) ExecX(ctx context.Context) {
	if err := mwicb.Exec(ctx); err != nil {
		panic(err)
	}
}
