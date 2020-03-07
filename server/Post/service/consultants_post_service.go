package service

import (
	"github.com/codeNight/server/Post"
	"github.com/codeNight/server/entity"
)

// PatientService implements Registeration.PatientService interface
type PostService struct {
	postRepo Post.PostRepository
}

// NewPatientService  returns a new PatientService object
func NewPostService(postRepository Post.PostRepository) Post.PostService {
	return &PostService{postRepo: postRepository}
}

// Patientes returns all stored application Patientes
func (ps *PostService) Posts() ([]entity.Post, []error) {
	posts, errs := ps.postRepo.Posts()
	if len(errs) > 0 {
		return nil, errs
	}
	return posts, errs
}

// Patient retrieves an application Patient by its id
func (ps *PostService) Post(id uint) (*entity.Post, []error) {
	pst, errs := ps.postRepo.Post(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return pst, errs
}

// UpdatePatient updates  a given application Patient
func (ps *PostService) UpdatePost(post *entity.Post) (*entity.Post, []error) {
	pst, errs := ps.postRepo.UpdatePost(post)
	if len(errs) > 0 {
		return nil, errs
	}
	return pst, errs
}

// DeletePatient deletes a given application Patient
func (ps *PostService) DeletePost(id uint) (*entity.Post, []error) {
	pst, errs := ps.postRepo.DeletePost(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return pst, errs
}

// StoreUser stores a given application user
func (ps *PostService) StorePost(post *entity.Post) (*entity.Post, []error) {
	pst, errs := ps.postRepo.StorePost(post)
	if len(errs) > 0 {
		return nil, errs
	}
	return pst, errs
}
