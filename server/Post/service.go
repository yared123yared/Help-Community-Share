package Post
import "github.com/codeNight/server/entity"
type PostService interface {
	DeletePost(id uint) (*entity.Post, []error)
	UpdatePost(Post *entity.Post) (*entity.Post, []error)
	Posts() ([]entity.Post, []error)
	Post(id uint) (*entity.Post, []error)
	StorePost(post *entity.Post) (*entity.Post, []error)
}