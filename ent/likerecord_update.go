// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"memo/ent/likerecord"
	"memo/ent/memo"
	"memo/ent/predicate"
	"memo/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// LikeRecordUpdate is the builder for updating LikeRecord entities.
type LikeRecordUpdate struct {
	config
	hooks    []Hook
	mutation *LikeRecordMutation
}

// Where appends a list predicates to the LikeRecordUpdate builder.
func (lru *LikeRecordUpdate) Where(ps ...predicate.LikeRecord) *LikeRecordUpdate {
	lru.mutation.Where(ps...)
	return lru
}

// SetUserID sets the "user_id" field.
func (lru *LikeRecordUpdate) SetUserID(i int) *LikeRecordUpdate {
	lru.mutation.ResetUserID()
	lru.mutation.SetUserID(i)
	return lru
}

// AddUserID adds i to the "user_id" field.
func (lru *LikeRecordUpdate) AddUserID(i int) *LikeRecordUpdate {
	lru.mutation.AddUserID(i)
	return lru
}

// SetMemoID sets the "memo_id" field.
func (lru *LikeRecordUpdate) SetMemoID(u uuid.UUID) *LikeRecordUpdate {
	lru.mutation.SetMemoID(u)
	return lru
}

// SetCreatedAt sets the "created_at" field.
func (lru *LikeRecordUpdate) SetCreatedAt(t time.Time) *LikeRecordUpdate {
	lru.mutation.SetCreatedAt(t)
	return lru
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lru *LikeRecordUpdate) SetNillableCreatedAt(t *time.Time) *LikeRecordUpdate {
	if t != nil {
		lru.SetCreatedAt(*t)
	}
	return lru
}

// AddMemoIDs adds the "memos" edge to the Memo entity by IDs.
func (lru *LikeRecordUpdate) AddMemoIDs(ids ...uuid.UUID) *LikeRecordUpdate {
	lru.mutation.AddMemoIDs(ids...)
	return lru
}

// AddMemos adds the "memos" edges to the Memo entity.
func (lru *LikeRecordUpdate) AddMemos(m ...*Memo) *LikeRecordUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return lru.AddMemoIDs(ids...)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (lru *LikeRecordUpdate) AddUserIDs(ids ...int) *LikeRecordUpdate {
	lru.mutation.AddUserIDs(ids...)
	return lru
}

// AddUsers adds the "users" edges to the User entity.
func (lru *LikeRecordUpdate) AddUsers(u ...*User) *LikeRecordUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return lru.AddUserIDs(ids...)
}

// Mutation returns the LikeRecordMutation object of the builder.
func (lru *LikeRecordUpdate) Mutation() *LikeRecordMutation {
	return lru.mutation
}

// ClearMemos clears all "memos" edges to the Memo entity.
func (lru *LikeRecordUpdate) ClearMemos() *LikeRecordUpdate {
	lru.mutation.ClearMemos()
	return lru
}

// RemoveMemoIDs removes the "memos" edge to Memo entities by IDs.
func (lru *LikeRecordUpdate) RemoveMemoIDs(ids ...uuid.UUID) *LikeRecordUpdate {
	lru.mutation.RemoveMemoIDs(ids...)
	return lru
}

// RemoveMemos removes "memos" edges to Memo entities.
func (lru *LikeRecordUpdate) RemoveMemos(m ...*Memo) *LikeRecordUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return lru.RemoveMemoIDs(ids...)
}

// ClearUsers clears all "users" edges to the User entity.
func (lru *LikeRecordUpdate) ClearUsers() *LikeRecordUpdate {
	lru.mutation.ClearUsers()
	return lru
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (lru *LikeRecordUpdate) RemoveUserIDs(ids ...int) *LikeRecordUpdate {
	lru.mutation.RemoveUserIDs(ids...)
	return lru
}

// RemoveUsers removes "users" edges to User entities.
func (lru *LikeRecordUpdate) RemoveUsers(u ...*User) *LikeRecordUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return lru.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lru *LikeRecordUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, lru.sqlSave, lru.mutation, lru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lru *LikeRecordUpdate) SaveX(ctx context.Context) int {
	affected, err := lru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lru *LikeRecordUpdate) Exec(ctx context.Context) error {
	_, err := lru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lru *LikeRecordUpdate) ExecX(ctx context.Context) {
	if err := lru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lru *LikeRecordUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(likerecord.Table, likerecord.Columns, sqlgraph.NewFieldSpec(likerecord.FieldID, field.TypeInt))
	if ps := lru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lru.mutation.UserID(); ok {
		_spec.SetField(likerecord.FieldUserID, field.TypeInt, value)
	}
	if value, ok := lru.mutation.AddedUserID(); ok {
		_spec.AddField(likerecord.FieldUserID, field.TypeInt, value)
	}
	if value, ok := lru.mutation.MemoID(); ok {
		_spec.SetField(likerecord.FieldMemoID, field.TypeUUID, value)
	}
	if value, ok := lru.mutation.CreatedAt(); ok {
		_spec.SetField(likerecord.FieldCreatedAt, field.TypeTime, value)
	}
	if lru.mutation.MemosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   likerecord.MemosTable,
			Columns: []string{likerecord.MemosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memo.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lru.mutation.RemovedMemosIDs(); len(nodes) > 0 && !lru.mutation.MemosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   likerecord.MemosTable,
			Columns: []string{likerecord.MemosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memo.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lru.mutation.MemosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   likerecord.MemosTable,
			Columns: []string{likerecord.MemosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memo.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lru.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   likerecord.UsersTable,
			Columns: []string{likerecord.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lru.mutation.RemovedUsersIDs(); len(nodes) > 0 && !lru.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   likerecord.UsersTable,
			Columns: []string{likerecord.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lru.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   likerecord.UsersTable,
			Columns: []string{likerecord.UsersColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, lru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{likerecord.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lru.mutation.done = true
	return n, nil
}

// LikeRecordUpdateOne is the builder for updating a single LikeRecord entity.
type LikeRecordUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LikeRecordMutation
}

// SetUserID sets the "user_id" field.
func (lruo *LikeRecordUpdateOne) SetUserID(i int) *LikeRecordUpdateOne {
	lruo.mutation.ResetUserID()
	lruo.mutation.SetUserID(i)
	return lruo
}

// AddUserID adds i to the "user_id" field.
func (lruo *LikeRecordUpdateOne) AddUserID(i int) *LikeRecordUpdateOne {
	lruo.mutation.AddUserID(i)
	return lruo
}

// SetMemoID sets the "memo_id" field.
func (lruo *LikeRecordUpdateOne) SetMemoID(u uuid.UUID) *LikeRecordUpdateOne {
	lruo.mutation.SetMemoID(u)
	return lruo
}

// SetCreatedAt sets the "created_at" field.
func (lruo *LikeRecordUpdateOne) SetCreatedAt(t time.Time) *LikeRecordUpdateOne {
	lruo.mutation.SetCreatedAt(t)
	return lruo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lruo *LikeRecordUpdateOne) SetNillableCreatedAt(t *time.Time) *LikeRecordUpdateOne {
	if t != nil {
		lruo.SetCreatedAt(*t)
	}
	return lruo
}

// AddMemoIDs adds the "memos" edge to the Memo entity by IDs.
func (lruo *LikeRecordUpdateOne) AddMemoIDs(ids ...uuid.UUID) *LikeRecordUpdateOne {
	lruo.mutation.AddMemoIDs(ids...)
	return lruo
}

// AddMemos adds the "memos" edges to the Memo entity.
func (lruo *LikeRecordUpdateOne) AddMemos(m ...*Memo) *LikeRecordUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return lruo.AddMemoIDs(ids...)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (lruo *LikeRecordUpdateOne) AddUserIDs(ids ...int) *LikeRecordUpdateOne {
	lruo.mutation.AddUserIDs(ids...)
	return lruo
}

// AddUsers adds the "users" edges to the User entity.
func (lruo *LikeRecordUpdateOne) AddUsers(u ...*User) *LikeRecordUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return lruo.AddUserIDs(ids...)
}

// Mutation returns the LikeRecordMutation object of the builder.
func (lruo *LikeRecordUpdateOne) Mutation() *LikeRecordMutation {
	return lruo.mutation
}

// ClearMemos clears all "memos" edges to the Memo entity.
func (lruo *LikeRecordUpdateOne) ClearMemos() *LikeRecordUpdateOne {
	lruo.mutation.ClearMemos()
	return lruo
}

// RemoveMemoIDs removes the "memos" edge to Memo entities by IDs.
func (lruo *LikeRecordUpdateOne) RemoveMemoIDs(ids ...uuid.UUID) *LikeRecordUpdateOne {
	lruo.mutation.RemoveMemoIDs(ids...)
	return lruo
}

// RemoveMemos removes "memos" edges to Memo entities.
func (lruo *LikeRecordUpdateOne) RemoveMemos(m ...*Memo) *LikeRecordUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return lruo.RemoveMemoIDs(ids...)
}

// ClearUsers clears all "users" edges to the User entity.
func (lruo *LikeRecordUpdateOne) ClearUsers() *LikeRecordUpdateOne {
	lruo.mutation.ClearUsers()
	return lruo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (lruo *LikeRecordUpdateOne) RemoveUserIDs(ids ...int) *LikeRecordUpdateOne {
	lruo.mutation.RemoveUserIDs(ids...)
	return lruo
}

// RemoveUsers removes "users" edges to User entities.
func (lruo *LikeRecordUpdateOne) RemoveUsers(u ...*User) *LikeRecordUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return lruo.RemoveUserIDs(ids...)
}

// Where appends a list predicates to the LikeRecordUpdate builder.
func (lruo *LikeRecordUpdateOne) Where(ps ...predicate.LikeRecord) *LikeRecordUpdateOne {
	lruo.mutation.Where(ps...)
	return lruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (lruo *LikeRecordUpdateOne) Select(field string, fields ...string) *LikeRecordUpdateOne {
	lruo.fields = append([]string{field}, fields...)
	return lruo
}

// Save executes the query and returns the updated LikeRecord entity.
func (lruo *LikeRecordUpdateOne) Save(ctx context.Context) (*LikeRecord, error) {
	return withHooks(ctx, lruo.sqlSave, lruo.mutation, lruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lruo *LikeRecordUpdateOne) SaveX(ctx context.Context) *LikeRecord {
	node, err := lruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (lruo *LikeRecordUpdateOne) Exec(ctx context.Context) error {
	_, err := lruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lruo *LikeRecordUpdateOne) ExecX(ctx context.Context) {
	if err := lruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lruo *LikeRecordUpdateOne) sqlSave(ctx context.Context) (_node *LikeRecord, err error) {
	_spec := sqlgraph.NewUpdateSpec(likerecord.Table, likerecord.Columns, sqlgraph.NewFieldSpec(likerecord.FieldID, field.TypeInt))
	id, ok := lruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "LikeRecord.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := lruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, likerecord.FieldID)
		for _, f := range fields {
			if !likerecord.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != likerecord.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := lruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lruo.mutation.UserID(); ok {
		_spec.SetField(likerecord.FieldUserID, field.TypeInt, value)
	}
	if value, ok := lruo.mutation.AddedUserID(); ok {
		_spec.AddField(likerecord.FieldUserID, field.TypeInt, value)
	}
	if value, ok := lruo.mutation.MemoID(); ok {
		_spec.SetField(likerecord.FieldMemoID, field.TypeUUID, value)
	}
	if value, ok := lruo.mutation.CreatedAt(); ok {
		_spec.SetField(likerecord.FieldCreatedAt, field.TypeTime, value)
	}
	if lruo.mutation.MemosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   likerecord.MemosTable,
			Columns: []string{likerecord.MemosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memo.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lruo.mutation.RemovedMemosIDs(); len(nodes) > 0 && !lruo.mutation.MemosCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   likerecord.MemosTable,
			Columns: []string{likerecord.MemosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memo.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lruo.mutation.MemosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   likerecord.MemosTable,
			Columns: []string{likerecord.MemosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memo.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lruo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   likerecord.UsersTable,
			Columns: []string{likerecord.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lruo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !lruo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   likerecord.UsersTable,
			Columns: []string{likerecord.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lruo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   likerecord.UsersTable,
			Columns: []string{likerecord.UsersColumn},
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
	_node = &LikeRecord{config: lruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, lruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{likerecord.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	lruo.mutation.done = true
	return _node, nil
}
