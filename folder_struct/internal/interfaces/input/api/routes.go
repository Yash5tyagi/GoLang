package api

import (
	"krriya/internal/config"
	"krriya/internal/usecase/auth"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type ApiServer struct {
	Srv  *chi.Mux
	Port string
}

func SetUpRoutes(Srv *config.Server) *chi.Mux {
	allowedOrigins := []string{}

	_ = auth.NewService()
	//Initialise the router
	r := chi.NewRouter()

	//Add cors to the server
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "Set-Cookie", "Cookie"},
		AllowCredentials: true,
		ExposedHeaders:   []string{},
	}))

	r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running"))
	})

	r.Route("/api", func(r chi.Router) {

	})

	return r
}
