package internal

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
	"github.com/sanokkk/go_auth/internal/config"
	"github.com/sanokkk/go_auth/internal/db/repo"
)

type ApiConfig struct {
	DB *repo.Queries
}

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

	fmt.Println(config.DbURL)
	conn, err := sql.Open("postgres", config.DbURL)
	if err != nil {
		log.Fatal("error while open sql connection: ", err.Error())
	}

	apiCfg := ApiConfig{
		DB: repo.New(conn),
	}

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
		w.WriteHeader(200)
	})

	router1 := chi.NewRouter()
	router1.Post("/register", apiCfg.handlreCreateUser)
	router1.Post("/login", apiCfg.handleLogin)
	router1.Get("/welcome", apiCfg.handleAuth(apiCfg.handleWelcome))
	router1.Post("/reauth", apiCfg.handleReauth)

	router.Mount("/auth", router1)

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(config.PORT),
		Handler: router,
	}

	fmt.Println("running server on port ", config.PORT)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("error while running server")
	}
}
