package signUp

import "github.com/codeNight/server/entity"

// UserRepository specifies application user related database operations
type UserService interface {
	DeleteUser(id uint) (*entity.User, []error)
	UpdateUser(user *entity.User) (*entity.User, []error)
	Users() ([]entity.User, []error)
	User(id uint) (*entity.User, []error)
	StoreUser(user *entity.User) (*entity.User, []error)
}
