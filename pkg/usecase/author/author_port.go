package author

type AuthorInputPort interface {
	FindAll(interface{}) (interface{}, error)
	Create(interface{}) (interface{}, error)
}

type AuthorOutputPort interface {
	GetAuthorResponse(interface{}) (interface{}, error)
}
