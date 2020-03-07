package Post_Handler



import (
	"encoding/json"
	"fmt"
	_ "fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/codeNight/server/Post"
	"github.com/codeNight/server/entity"
)

type PostHandler struct {
	postService Post.PostService
}

// NewAdminCommentHandler returns new AdminCommentHandler object
func NewPostHandler(pstService Post.PostService) *PostHandler {
	return &PostHandler{postService: pstService}
}

// GetUsers handles GET /v1/admin/users request
func (ph *PostHandler) GetPosts(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	posts, errs := ph.postService.Posts()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(posts, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

// GetSingleUsers handles GET /v1/admin/users/:id request
func (ph *PostHandler) GetSinglePost(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {
	fmt.Println(" i am about to get single value")

	id, err := strconv.Atoi(ps.ByName("id"))
	fmt.Println(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	post, errs := ph.postService.Post(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(post, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// PostUser handles POST /v1/admin/users request
func (ph *PostHandler) PostPosts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println(" i am at the post method")

	l := r.ContentLength
	body := make([]byte, l)
	fmt.Println(" ia have changed the data to byte")
	fmt.Println(string(body))
	r.Body.Read(body)
	post := &entity.Post{}

	err := json.Unmarshal(body, post)
	fmt.Println("thise is the unmarchal jeson")
	fmt.Println(post)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	post, errs := ph.postService.StorePost(post)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("/v1/admin/users/%d")
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}

// PutComment handles PUT /v1/admin/comments/:id request
func (ph *PostHandler) PutPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
fmt.Println("this is the putUser method")
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	post, errs := ph.postService.Post(uint(id))
fmt.Println("this is the user")
fmt.Println(post)
	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)
	fmt.Println("this is the data before the changed to the jeson")
	fmt.Println(string(body))
	json.Unmarshal(body, &post)
	fmt.Println("this is the updated data")
	fmt.Println(post)
	post, errs = ph.postService.UpdatePost(post)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(post, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// DeleteComment handles DELETE /v1/admin/comments/:id request
func (ph *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println(" i amabout to delete ")
	id, err := strconv.Atoi(ps.ByName("id"))
	fmt.Println(id)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := ph.postService.DeletePost(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
