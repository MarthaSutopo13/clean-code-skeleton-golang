package service

import (
	"basesvc_v2/pkg/domain/entity"
	"basesvc_v2/pkg/usecase/author"
	"context"
)

type AuthorService struct {
	repoAuthor author.AuthorInputPort
}

func NewAuthorService(e author.AuthorInputPort) *AuthorService {
	return &AuthorService{
		repoAuthor: e,
	}
}

func (a *AuthorService) Create(ctx context.Context, req *entity.Author) (interface{}, error) {
	res, err := a.repoAuthor.Create(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
