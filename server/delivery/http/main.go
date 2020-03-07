package main

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"

	"github.com/codeNight/server/delivery/http/handler/signUp_Handler"
	"github.com/codeNight/server/delivery/http/handler/Post_Handler"
	"github.com/codeNight/server/signUp/repository"
	post_repository "github.com/codeNight/server/Post/repository"
	post_service "github.com/codeNight/server/Post/service"
	"github.com/codeNight/server/signUp/service"
)

func main() {
	dbconn, err := gorm.Open("postgres", "postgres://postgres:P@$$w0rDd@localhost/council?sslmode=disable")

	if err != nil {
		panic(err)
	}
	router := httprouter.New()
	fmt.Println("hi")

	
	// user registeration path registeration
	userRepo := repository.NewUserGormRepo(dbconn)
	userSrv := service.NewUserService(userRepo)
	userHandler :=signUp_Handler.NewUserHandler(userSrv)

	router.GET("/v1/user/users/:id", userHandler.GetSingleUser)
	router.GET("/v1/user/users/", userHandler.GetUsers)
	router.PUT("/v1/user/users/:id", userHandler.PutUser)
	router.POST("/v1/user/users/", userHandler.PostUser)
	// consultants posts path registeration
	postRepo := post_repository.NewPostGormRepo(dbconn)
	postSrv := post_service.NewPostService(postRepo)
	postHandler :=Post_Handler.NewPostHandler(postSrv)

	router.GET("/v1/consultants/posts/:id", postHandler.GetSinglePost)
	router.GET("/v1/consultants/posts/", postHandler.GetPosts)
	router.PUT("/v1/consultants/posts/:id", postHandler.PutPost)
	router.POST("/v1/consultants/posts/", postHandler.PostPosts)

	
	//
	http.ListenAndServe(":8180", router)

}
