package router

import (
	"friend_management/cmd/controller"
	"friend_management/internal/db"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

// Handler returns the http handler that handles all requests
func Handler(db db.Executor) http.Handler {
	r := chi.NewRouter()

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
		userRouter.Get("/", controller.GetUser(db))
		userRouter.Get("/GetAll", controller.GetAllUsers(db))
		userRouter.Post("/connect", controller.ConnectFriends(db))
		userRouter.Get("/list", controller.FriendList(db))
		userRouter.Post("/common", controller.CommonFriends(db))
		userRouter.Post("/subscribe", controller.Subscription(db))
		userRouter.Post("/blocked", controller.Blocked(db))
		userRouter.Post("/send", controller.SendUpdate(db))
	})

	return r
}
