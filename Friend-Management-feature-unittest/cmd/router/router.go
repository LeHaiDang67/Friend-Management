package router

import (
	"database/sql"
	"friend_management/cmd/controller"
	"friend_management/intenal/services"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

// Handler returns the http handler that handles all requests
func Handler(dbconn *sql.DB) http.Handler {
	r := chi.NewRouter()
	st := services.NewManager(dbconn)
	// Basic CORS
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-Id-Token"},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	r.Route("/friend", func(userRouter chi.Router) {
		userRouter.Get("/", controller.GetUser(st))
		userRouter.Get("/GetAll", controller.GetAllUsers(st))
		userRouter.Post("/connect", controller.ConnectFriends(st))
		userRouter.Get("/list", controller.FriendList(st))
		userRouter.Post("/common", controller.CommonFriends(st))
		userRouter.Post("/subscribe", controller.Subscription(st))
		userRouter.Post("/blocked", controller.Blocked(st))
		userRouter.Post("/send", controller.SendUpdate(st))
	})

	return r
}
