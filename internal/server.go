package internal

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/sanokkk/go_auth/internal/config"
)

func Serve() {
	config := config.GetConfig()

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT"},
		AllowCredentials: false,
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
	}))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
		w.WriteHeader(200)
	})

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(config.PORT),
		Handler: router,
	}

	fmt.Println("running server on port ", config.PORT)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("error while running server")
	}
}
