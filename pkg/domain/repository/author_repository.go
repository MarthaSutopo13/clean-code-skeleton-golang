package repository

type AuthorRepository interface {
	FindAll(interface{}) (interface{}, error)
	Create(interface{}) (interface{}, error)
}
