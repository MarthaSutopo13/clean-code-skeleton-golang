package author

import "basesvc_v2/pkg/domain/repository"

type AuthorUsecase struct {
	repo repository.AuthorRepository
	out  AuthorOutputPort
}

//NewAuthorUsecase
func NewAuthorUsecase(re repository.AuthorRepository, o AuthorOutputPort) *AuthorUsecase {
	return &AuthorUsecase{
		repo: re,
		out:  o,
	}
}

//CreateAuthor
func (aut *AuthorUsecase) Create(au interface{}) (interface{}, error) {
	res, err := aut.repo.Create(au)
	if err != nil {
		return nil, err
	}

	return aut.out.GetAuthorResponse(res)
}

//Find All
func (aut *AuthorUsecase) FindAll(au interface{}) (interface{}, error) {
	res, err := aut.repo.FindAll(au)
	if err != nil {
		return nil, err
	}

	return aut.out.GetAuthorResponse(res)
}
