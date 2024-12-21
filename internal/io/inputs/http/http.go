package HTTPHandler

import (
	"log"
	"net/http"

	"github.com/JaredSnapp/go_backend/internal/models"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"github.com/swaggest/openapi-go/openapi31"
	"github.com/swaggest/rest/web"
	swgui "github.com/swaggest/swgui/v5emb"
)

func (h Handler) Serve() {
	r := h.NewRouter()

	log.Println("Starting service at http://localhost:3000")
	if err := http.ListenAndServe("localhost:3000", r); err != nil {
		log.Fatal(err)
	}
}

func (h Handler) NewRouter() *web.Service {
	// Service initializes router with required middlewares.
	service := web.NewService(openapi31.NewReflector())

	// It allows OpenAPI configuration.
	service.OpenAPISchema().SetTitle("Go Backend API")
	service.OpenAPISchema().SetDescription("This service provides backend API in Go.")
	service.OpenAPISchema().SetVersion("v1.0.0")

	// c := cors.New(cors.Options{
	// 	AllowedOrigins: h.corsAllowedOrigin,
	// 	AllowedMethods: []string{"GET", "PUT", "POST", "PATCH", "DELETE", "HEAD", "OPTIONS"},
	// 	AllowedHeaders: []string{"*"},
	// 	Debug:          true,
	// })

	// service.Use(c.Handler)

	// Additional middlewares can be added.
	service.Use(
		middleware.StripSlashes,
		middleware.Logger,

		cors.AllowAll().Handler, // "github.com/rs/cors", 3rd-party CORS middleware can also be configured here.
	)

	service.Get("/persons", h.getPersons())
	service.Post("/persons", h.postPerson())
	service.Put("/persons/{id}", h.putPerson())
	service.Delete("/persons/{id}", h.deletePerson())

	goals := NewCrud[models.GoalMetaData, models.GoalPutInput]("goal", h.GoalsService)

	service.Get("/goal", goals.Get())
	service.Post("/goal", goals.Post())
	service.Put("/goal/{id}", goals.Put())
	service.Delete("/goal/{id}", goals.Delete())

	service.Docs("/docs", swgui.New)

	// Use cases can be mounted using short syntax .<Method>(...).
	// service.Post("/albums", postAlbums(), nethttp.SuccessStatus(http.StatusCreated))

	return service
}
