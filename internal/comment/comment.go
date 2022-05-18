package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrorMessage        = errors.New("an error ocurred")
	ErrorNotImplemented = errors.New("not implemented")
)

type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

type Service struct {
	Store Store
}

func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retriving a comment")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrorMessage
	}
	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return ErrorNotImplemented
}

func (*Service) DeleteCommment(ctx context.Context, id string) error {
	return ErrorNotImplemented
}

func (*Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, ErrorNotImplemented
}
