package main

import (
	"os"

	"github.com/Amirhos-esm/portfolio/graph"
	"github.com/Amirhos-esm/portfolio/models"
	"github.com/Amirhos-esm/portfolio/repository"
	"github.com/Amirhos-esm/portfolio/repository/json"
	"github.com/Amirhos-esm/portfolio/repository/sqlite"
	"github.com/Amirhos-esm/portfolio/util"
	"github.com/vektah/gqlparser/v2/ast"

	"github.com/gin-gonic/gin"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
)

type Application struct {
	repo             repository.Repository
	messageRepo      repository.MessageRepository
	auth             Auth
	host             string
	password         string
	enablePlayground bool
	staticFolder     string
}

func main() {
	app := NewApplication()

	router := gin.Default()
	app.registerRoutes(router)

	router.Run(app.host)
}

func NewApplication() *Application {

	password := os.Getenv("PASS")
	if password == "" {
		password = "demo"
	}
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost:8080"
	}

	enablePlayground := os.Getenv("PLAYGROUND") != ""

	repo, err := json.NewJSONRepository(
		"data.json",
		util.Ptr(models.GetMockData()),
	)
	if err != nil {
		panic(err)
	}

	messageRepo, err := sqlite.NewMessageGorm()
	if err != nil {
		panic(err)
	}

	return &Application{
		repo:             repo,
		messageRepo: messageRepo,
		auth:             NewAuth(),
		password:         password,
		enablePlayground: !enablePlayground,
		staticFolder:     "./static",
		host:             host,
	}
}

func (app *Application) newGraphQLServer() gin.HandlerFunc {

	srv := handler.New(
		graph.NewExecutableSchema(
			graph.Config{Resolvers: &Resolver{app}},
		),
	)

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})

	srv.Use(extension.AutomaticPersistedQuery{Cache: lru.New[string](100)})

	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

func (app *Application) registerRoutes(r *gin.Engine) {

	r.Use(CORSMiddleware())
	if app.enablePlayground {
		r.GET("/GraphQL", gin.WrapH(
			playground.Handler("GraphQL", "/api/graphql"),
		))
	}
	//init admin panel
	initAdminPanel(r)
	// public routes
	r.GET("/", app.LandingPageHandler)
	r.GET("/projects/:id", app.ProjectHandler)
	r.Static("/static", app.staticFolder)
	r.POST("api/login", app.loginHandler)
	//create message
	r.POST("api/message", app.createMessageHandler)

	// protected routes
	// admin := app.AuthRequiredMiddleware()
	admin := func(*gin.Context) {}

	r.POST("/api/graphql", admin, app.newGraphQLServer())

	r.POST("api/projects/:id/gallery", admin, app.addProjectGallery)
	r.POST("api/profile", admin, app.profileUploadHandler)
	r.POST("api/resume", admin, app.resumeUploadHandler)
}
