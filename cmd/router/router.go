package router

import (
	"database/sql"
	"friend_management/cmd/controller"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

// Handler returns the http handler that handles all requests
func Handler(db *sql.DB) http.Handler {
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
		userRouter.Put("/", controller.UpdateUser(db))
		userRouter.Post("/connect", controller.ConnectFriends(db))
		userRouter.Get("/list", controller.FriendList(db))
		userRouter.Post("/common", controller.CommonFriends(db))
		userRouter.Post("/subscribe", controller.Subscription(db))
		userRouter.Post("/blocked", controller.Blocked(db))
	})

	return r
}
