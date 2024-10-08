package httpserver

import (
	"fmt"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"kanban/internal/config"
	projectservice "kanban/internal/domain/usecases/services/project_service"
	taskservice "kanban/internal/domain/usecases/services/task_service"
	userservice "kanban/internal/domain/usecases/services/user_service"
	"kanban/internal/http-server/handlers"
	"kanban/internal/http-server/middleware/authorization"
	"kanban/internal/http-server/middleware/cors"
	"kanban/internal/http-server/middleware/logger"
	"log/slog"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "kanban/docs"
)

// @title Kaban API
// @version 1.0
// @BasePath /api/v1
func New(
	log *slog.Logger,
	cfg *config.Config,
	taskService taskservice.TaskService,
	userService userservice.UserService,
	projectsService projectservice.ProjectService,
) *chi.Mux {
	h := handlers.NewHandler(
		log,
		taskService,
		userService,
		projectsService,
	)

	router := chi.NewRouter()

	router.Use(middleware.RequestID)    // Добавление RequestID
	router.Use(middleware.Recoverer)    // Обработка паник
	router.Use(middleware.URLFormat)    // Обработка URL
	router.Use(middleware.CleanPath)    // Удаление лишних слешей
	router.Use(middleware.StripSlashes) // Удаление последнего слеша
	router.Use(logger.New(log))

	router.Use(cors.New(cors.CORSOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		ExposedHeaders:   []string{"X-Custom-Header"},
		AllowCredentials: true,
		MaxAge:           3600,
	}))

	router.Group(func(r chi.Router) {
		r.Use(middleware.BasicAuth("Swagger Protected Area", map[string]string{
			cfg.Swag.Login: cfg.Swag.Password,
		}))

		r.Get("/swagger/*", httpSwagger.Handler(
			httpSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", cfg.Swag.Endpoint)),
		))
	})

	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/info", h.Info)

		r.Route("/user", func(r chi.Router) {
			r.Post("/register", h.Register)
			r.Post("/login", h.Login)
			r.Post("/logout", h.Logout)
			r.Get("/me", h.Me)
		})

		r.Group(func(r chi.Router) {
			r.Use(authorization.AuthorizationMiddleware(cfg.Auth.Secret))

			r.Route("/projects", func(r chi.Router) {
				r.Post("/", h.CreateProject)
				r.Put("/", h.UpdateProject)
				r.Delete("/{id}", h.DeleteProject)
				r.Get("/{id}", h.GetProject)
				r.Get("/{name}", h.GetProjects)
			})
		})
	})

	return router
}
