package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrNotImplemented  = errors.New("not implemented")
)

// Comment - a representation of the comment structure for server
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store - this interface defines all of the methods
// that our service need in order to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
	PostComment(context.Context, Comment) (Comment, error)
}

// Service - the struct for our comment service
type Service struct {
	Store Store
}

// NewService - returns a new comment service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

// GetComment - retrieves comments by their ID from the database
func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retrieving the comment")
	ctx = context.WithValue(ctx, "request_id", "unique-string")
	fmt.Println((ctx.Value("request_id")))
	return Comment{}, nil
}

// UpdateComment - updates a comment by ID with new comment info
func (s *Service) UpdateComment(ctx context.Context, ID string, newCmt Comment) error {
	return ErrNotImplemented
}

// DeleteComment - deletes a comment from the database by ID
func (s *Service) DeleteComment(ctx context.Context, ID string) error {
	return ErrNotImplemented
}

// PostComment - adds a new comment to the database
func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {
	postedComment, err := s.Store.PostComment(ctx, cmt)
	if err != nil {
		return Comment{}, err
	}
	return postedComment, nil
}
