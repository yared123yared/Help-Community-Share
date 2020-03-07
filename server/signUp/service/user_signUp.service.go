package service

import (
	"github.com/codeNight/server/signUp"
	"github.com/codeNight/server/entity"
)

// PatientService implements Registeration.PatientService interface
type UserService struct {
	userRepo signUp.UserRepository
}

// NewPatientService  returns a new PatientService object
func NewUserService(userRepository signUp.UserRepository) signUp.UserService {
	return &UserService{userRepo: userRepository}
}

// Patientes returns all stored application Patientes
func (us *UserService) Users() ([]entity.User, []error) {
	petientes, errs := us.userRepo.Users()
	if len(errs) > 0 {
		return nil, errs
	}
	return petientes, errs
}

// Patient retrieves an application Patient by its id
func (us *UserService) User(id uint) (*entity.User, []error) {
	pst, errs := us.userRepo.User(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return pst, errs
}

// UpdatePatient updates  a given application Patient
func (us *UserService) UpdateUser(user *entity.User) (*entity.User, []error) {
	pst, errs := us.userRepo.UpdateUser(user)
	if len(errs) > 0 {
		return nil, errs
	}
	return pst, errs
}

// DeletePatient deletes a given application Patient
func (us *UserService) DeleteUser(id uint) (*entity.User, []error) {
	pst, errs := us.userRepo.DeleteUser(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return pst, errs
}

// StoreUser stores a given application user
func (us *UserService) StoreUser(user *entity.User) (*entity.User, []error) {
	pst, errs := us.userRepo.StoreUser(user)
	if len(errs) > 0 {
		return nil, errs
	}
	return pst, errs
}
