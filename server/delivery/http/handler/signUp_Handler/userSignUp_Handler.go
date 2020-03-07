package signUp_Handler

import (
	"encoding/json"
	"fmt"
	_ "fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/codeNight/server/signUp"
	"github.com/codeNight/server/entity"
)

type UserHandler struct {
	userService signUp.UserService
}

// NewAdminCommentHandler returns new AdminCommentHandler object
func NewUserHandler(usrService signUp.UserService) *UserHandler {
	return &UserHandler{userService: usrService}
}

// GetUsers handles GET /v1/admin/users request
func (uh *UserHandler) GetUsers(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	users, errs := uh.userService.Users()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(users, "", "\t\t")

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
func (uh *UserHandler) GetSingleUser(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {
	fmt.Println(" i am about to get single value")

	id, err := strconv.Atoi(ps.ByName("id"))
	fmt.Println(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	user, errs := uh.userService.User(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(user, "", "\t\t")

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
func (uh *UserHandler) PostUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println(" i am at the post method")

	l := r.ContentLength
	body := make([]byte, l)
	fmt.Println(" ia have changed the data to byte")
	fmt.Println(string(body))
	r.Body.Read(body)
	user := &entity.User{}

	err := json.Unmarshal(body, user)
	fmt.Println("thise is the unmarchal jeson")
	fmt.Println(user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	user, errs := uh.userService.StoreUser(user)

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
func (uh *UserHandler) PutUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
fmt.Println("this is the putUser method")
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	user, errs := uh.userService.User(uint(id))
fmt.Println("this is the user")
fmt.Println(user)
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
	json.Unmarshal(body, &user)
	fmt.Println("this is the updated data")
	fmt.Println(user)
	user, errs = uh.userService.UpdateUser(user)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(user, "", "\t\t")

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
func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println(" i amabout to delete ")
	id, err := strconv.Atoi(ps.ByName("id"))
	fmt.Println(id)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := uh.userService.DeleteUser(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
