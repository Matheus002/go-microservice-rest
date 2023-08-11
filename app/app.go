package app

import (
	"net/http"

	"github.com/Matheus002/go-microservice-rest/internal/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":3333", nil); err != nil {
		panic(err)
	}
}
