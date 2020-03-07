package Post

import "github.com/codeNight/server/entity"

// UserRepository specifies application user related database operations
type PostRepository interface {
	DeletePost(id uint) (*entity.Post, []error)
	UpdatePost(Post *entity.Post) (*entity.Post, []error)
	Posts() ([]entity.Post, []error)
	Post(id uint) (*entity.Post, []error)
	StorePost(post *entity.Post) (*entity.Post, []error)
}
