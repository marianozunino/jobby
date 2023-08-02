// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/marianozunino/cc-backend-go/ent/post"
	"github.com/marianozunino/cc-backend-go/ent/postcategory"
	"github.com/marianozunino/cc-backend-go/ent/user"
)

// PostCreate is the builder for creating a Post entity.
type PostCreate struct {
	config
	mutation *PostMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (pc *PostCreate) SetTitle(s string) *PostCreate {
	pc.mutation.SetTitle(s)
	return pc
}

// SetContent sets the "content" field.
func (pc *PostCreate) SetContent(s string) *PostCreate {
	pc.mutation.SetContent(s)
	return pc
}

// SetSlug sets the "slug" field.
func (pc *PostCreate) SetSlug(s string) *PostCreate {
	pc.mutation.SetSlug(s)
	return pc
}

// SetIsHighlighted sets the "is_highlighted" field.
func (pc *PostCreate) SetIsHighlighted(b bool) *PostCreate {
	pc.mutation.SetIsHighlighted(b)
	return pc
}

// SetIsPublished sets the "is_published" field.
func (pc *PostCreate) SetIsPublished(b bool) *PostCreate {
	pc.mutation.SetIsPublished(b)
	return pc
}

// SetPublishedAt sets the "published_at" field.
func (pc *PostCreate) SetPublishedAt(t time.Time) *PostCreate {
	pc.mutation.SetPublishedAt(t)
	return pc
}

// SetNillablePublishedAt sets the "published_at" field if the given value is not nil.
func (pc *PostCreate) SetNillablePublishedAt(t *time.Time) *PostCreate {
	if t != nil {
		pc.SetPublishedAt(*t)
	}
	return pc
}

// SetPreviewImage sets the "preview_image" field.
func (pc *PostCreate) SetPreviewImage(s string) *PostCreate {
	pc.mutation.SetPreviewImage(s)
	return pc
}

// SetNillablePreviewImage sets the "preview_image" field if the given value is not nil.
func (pc *PostCreate) SetNillablePreviewImage(s *string) *PostCreate {
	if s != nil {
		pc.SetPreviewImage(*s)
	}
	return pc
}

// SetAuthorID sets the "author_id" field.
func (pc *PostCreate) SetAuthorID(u uuid.UUID) *PostCreate {
	pc.mutation.SetAuthorID(u)
	return pc
}

// SetNillableAuthorID sets the "author_id" field if the given value is not nil.
func (pc *PostCreate) SetNillableAuthorID(u *uuid.UUID) *PostCreate {
	if u != nil {
		pc.SetAuthorID(*u)
	}
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *PostCreate) SetCreatedAt(t time.Time) *PostCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PostCreate) SetNillableCreatedAt(t *time.Time) *PostCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PostCreate) SetUpdatedAt(t time.Time) *PostCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PostCreate) SetNillableUpdatedAt(t *time.Time) *PostCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetDeletedAt sets the "deleted_at" field.
func (pc *PostCreate) SetDeletedAt(t time.Time) *PostCreate {
	pc.mutation.SetDeletedAt(t)
	return pc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pc *PostCreate) SetNillableDeletedAt(t *time.Time) *PostCreate {
	if t != nil {
		pc.SetDeletedAt(*t)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PostCreate) SetID(u uuid.UUID) *PostCreate {
	pc.mutation.SetID(u)
	return pc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (pc *PostCreate) SetUserID(id uuid.UUID) *PostCreate {
	pc.mutation.SetUserID(id)
	return pc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (pc *PostCreate) SetNillableUserID(id *uuid.UUID) *PostCreate {
	if id != nil {
		pc = pc.SetUserID(*id)
	}
	return pc
}

// SetUser sets the "user" edge to the User entity.
func (pc *PostCreate) SetUser(u *User) *PostCreate {
	return pc.SetUserID(u.ID)
}

// AddPostCategoryIDs adds the "post_category" edge to the PostCategory entity by IDs.
func (pc *PostCreate) AddPostCategoryIDs(ids ...uuid.UUID) *PostCreate {
	pc.mutation.AddPostCategoryIDs(ids...)
	return pc
}

// AddPostCategory adds the "post_category" edges to the PostCategory entity.
func (pc *PostCreate) AddPostCategory(p ...*PostCategory) *PostCreate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddPostCategoryIDs(ids...)
}

// Mutation returns the PostMutation object of the builder.
func (pc *PostCreate) Mutation() *PostMutation {
	return pc.mutation
}

// Save creates the Post in the database.
func (pc *PostCreate) Save(ctx context.Context) (*Post, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PostCreate) SaveX(ctx context.Context) *Post {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PostCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PostCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PostCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := post.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := post.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PostCreate) check() error {
	if _, ok := pc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Post.title"`)}
	}
	if _, ok := pc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Post.content"`)}
	}
	if _, ok := pc.mutation.Slug(); !ok {
		return &ValidationError{Name: "slug", err: errors.New(`ent: missing required field "Post.slug"`)}
	}
	if _, ok := pc.mutation.IsHighlighted(); !ok {
		return &ValidationError{Name: "is_highlighted", err: errors.New(`ent: missing required field "Post.is_highlighted"`)}
	}
	if _, ok := pc.mutation.IsPublished(); !ok {
		return &ValidationError{Name: "is_published", err: errors.New(`ent: missing required field "Post.is_published"`)}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Post.created_at"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Post.updated_at"`)}
	}
	return nil
}

func (pc *PostCreate) sqlSave(ctx context.Context) (*Post, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PostCreate) createSpec() (*Post, *sqlgraph.CreateSpec) {
	var (
		_node = &Post{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(post.Table, sqlgraph.NewFieldSpec(post.FieldID, field.TypeUUID))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pc.mutation.Title(); ok {
		_spec.SetField(post.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := pc.mutation.Content(); ok {
		_spec.SetField(post.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := pc.mutation.Slug(); ok {
		_spec.SetField(post.FieldSlug, field.TypeString, value)
		_node.Slug = value
	}
	if value, ok := pc.mutation.IsHighlighted(); ok {
		_spec.SetField(post.FieldIsHighlighted, field.TypeBool, value)
		_node.IsHighlighted = value
	}
	if value, ok := pc.mutation.IsPublished(); ok {
		_spec.SetField(post.FieldIsPublished, field.TypeBool, value)
		_node.IsPublished = value
	}
	if value, ok := pc.mutation.PublishedAt(); ok {
		_spec.SetField(post.FieldPublishedAt, field.TypeTime, value)
		_node.PublishedAt = &value
	}
	if value, ok := pc.mutation.PreviewImage(); ok {
		_spec.SetField(post.FieldPreviewImage, field.TypeString, value)
		_node.PreviewImage = &value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(post.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(post.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.DeletedAt(); ok {
		_spec.SetField(post.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if nodes := pc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   post.UserTable,
			Columns: []string{post.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.AuthorID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PostCategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   post.PostCategoryTable,
			Columns: post.PostCategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(postcategory.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PostCreateBulk is the builder for creating many Post entities in bulk.
type PostCreateBulk struct {
	config
	builders []*PostCreate
}

// Save creates the Post entities in the database.
func (pcb *PostCreateBulk) Save(ctx context.Context) ([]*Post, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Post, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PostMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PostCreateBulk) SaveX(ctx context.Context) []*Post {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PostCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PostCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
