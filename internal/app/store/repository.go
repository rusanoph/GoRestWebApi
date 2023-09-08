package store

import "GoHttpRestApi/internal/app/model"

//UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
