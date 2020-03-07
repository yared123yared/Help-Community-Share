package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/codeNight/server/signUp"
	"github.com/codeNight/server/entity"
)

// UserGormRepo Implements the menu.UserRepository interface
type UserGormRepo struct {
	conn *gorm.DB
}

// NewUserGormRepo creates a new object of NewUserGormRepo
func NewUserGormRepo(db *gorm.DB) signUp.UserRepository {
	return &UserGormRepo{conn: db}
}

// Users return all users from the database
func (userRepo *UserGormRepo) Users() ([]entity.User, []error) {
	users := []entity.User{}
	errs := userRepo.conn.Find(&users).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return users, errs
}

// User retrieves a User by its id from the database
func (userRepo *UserGormRepo) User(id uint) (*entity.User, []error) {
	user := entity.User{}
	errs := userRepo.conn.First(&user, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &user, errs
}

// UpdateUser updates a given user in the database
func (userRepo *UserGormRepo) UpdateUser(user *entity.User) (*entity.User, []error) {
	usr := user
	errs := userRepo.conn.Save(usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// DeleteUser deletes a given usr from the database
func (userRepo *UserGormRepo) DeleteUser(id uint) (*entity.User, []error) {
	usr, errs := userRepo.User(id)
	if len(errs) > 0 {
		return nil, errs
	}
	
	errs = userRepo.conn.Delete(usr, id).GetErrors()
	
	return usr, errs
}

// StoreUser stores a new user into the database
func (userRepo *UserGormRepo) StoreUser(user *entity.User) (*entity.User, []error) {
	usr := user
	errs := userRepo.conn.Create(usr).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}
/*
func (userRepo *UserGormRepo) Profile(id uint) (*entity.User, []error) {
	user := entity.User{}
	errs := patientRepo.conn.First(&user, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &user, errs
}
*/
