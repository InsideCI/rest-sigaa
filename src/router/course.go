package router

import (
	"github.com/InsideCI/nego/src/controller/rest"
	"github.com/InsideCI/nego/src/driver"
	"github.com/InsideCI/nego/src/models"
	"github.com/InsideCI/nego/src/router/middlewares"
	"github.com/go-chi/chi"
)

func NewCourseRouter(db *driver.DB) func(r chi.Router) {
	generic := middlewares.GenericMiddleware{Type: models.Course{}}
	handlers := rest.NewCourseController(db)

	return func(r chi.Router) {
		// teachers
		r.With(generic.Persist).Post("/", handlers.Create)
		r.With(generic.Fetch).Get("/", handlers.Fetch)

		// teachers/{id}
		r.Route("/{id}", func(r chi.Router) {
			r.Use(generic.Identifier)
			r.Get("/", handlers.FetchOne)
			//r.With(generic.Persist).Put("/", handlers.Update)
			//r.Delete("/", handlers.Delete)
		})
	}
}
