package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/codeNight/server/Post"
	"github.com/codeNight/server/entity"
)

// UserGormRepo Implements the menu.UserRepository interface
type PostGormRepo struct {
	conn *gorm.DB
}

// NewUserGormRepo creates a new object of NewUserGormRepo
func NewPostGormRepo(db *gorm.DB) Post.PostRepository{
	return &PostGormRepo{conn: db}
}

// Users return all users from the database
func (postRepo *PostGormRepo) Posts() ([]entity.Post, []error) {
	posts := []entity.Post{}
	errs := postRepo.conn.Find(&posts).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return posts, errs
}

// User retrieves a User by its id from the database
func (postRepo *PostGormRepo) Post(id uint) (*entity.Post, []error) {
	post := entity.Post{}
	errs := postRepo.conn.First(&post, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &post, errs
}

// UpdateUser updates a given user in the database
func (postRepo *PostGormRepo) UpdatePost(post *entity.Post) (*entity.Post, []error) {
	pst := post
	errs := postRepo.conn.Save(pst).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return pst, errs
}

// DeleteUser deletes a given usr from the database
func (postRepo *PostGormRepo) DeletePost(id uint) (*entity.Post, []error) {
	pst, errs := postRepo.Post(id)
	if len(errs) > 0 {
		return nil, errs
	}
	
	errs = postRepo.conn.Delete(pst, id).GetErrors()
	
	return pst, errs
}

// StoreUser stores a new user into the database
func (postRepo *PostGormRepo) StorePost(post *entity.Post) (*entity.Post, []error) {
	pst := post
	errs := postRepo.conn.Create(pst).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return pst, errs
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
