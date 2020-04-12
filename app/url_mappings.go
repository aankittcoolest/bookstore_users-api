package app

import (
	"github.com/aankittcoolest/bookstore_users-api/controllers/ping"
	"github.com/aankittcoolest/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
}
