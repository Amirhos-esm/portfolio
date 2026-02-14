package main

import (
	"os"

	"github.com/Amirhos-esm/portfolio/graph"
	"github.com/Amirhos-esm/portfolio/models"
	"github.com/Amirhos-esm/portfolio/repository"
	"github.com/Amirhos-esm/portfolio/repository/json"
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
	auth             Auth
	password         string
	enablePlayground bool
	staticFolder     string
}

func main() {
	app := NewApplication()

	router := gin.Default()
	app.registerRoutes(router)

	router.Run("localhost:8080")
}

func NewApplication() *Application {

	password := os.Getenv("PASS")
	if password == "" {
		password = "demo"
	}

	enablePlayground := os.Getenv("PLAYGROUND") != ""

	repo, err := json.NewJSONRepository(
		"data.json",
		util.Ptr(models.GetMockData()),
	)
	if err != nil {
		panic(err)
	}

	return &Application{
		repo:             repo,
		auth:             NewAuth(),
		password:         password,
		enablePlayground: enablePlayground,
		staticFolder:     "./static",
	}
}

func (app *Application) newGraphQLServer() *handler.Server {

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

<<<<<<< HEAD
	srv.Use(extension.AutomaticPersistedQuery{Cache: lru.New[string](100)})
=======
	// 🔥 Mount GraphQL endpoint
	sv.POST("/graphql", gin.WrapH(srv))
	sv.GET("/graphql", gin.WrapH(srv))
	// Optional Playground
	sv.GET("/GraphQL", func(c *gin.Context) {
		playground.Handler("GraphQL", "/graphql").ServeHTTP(c.Writer, c.Request)
	})

	sv.GET("/", func(ctx *gin.Context) {
		data, err := app.repo.GetAllData()
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}
		views.Render(ctx.Writer, ctx.Request, pages.MainPage(data))
	})
	sv.POST("/login",func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"token":"asdadasd",
		})
	})
	sv.GET("/projects/:id", app.ProjectHandler)
	sv.POST("/projects/:id/gallery/", app.addProjectGallery)
	sv.POST("/profile", app.profileUploadHandler)
	sv.POST("/resume",app.resumeUploadHandler)
	sv.Static("/static", app.staticFolder)
	sv.Run("localhost:8080")
>>>>>>> c95efd425f98fb11522562f8fcef5dc3ec110aff

	return srv
}

func (app *Application) registerRoutes(r *gin.Engine) {

	srv := app.newGraphQLServer()

	if app.enablePlayground {
		r.GET("/GraphQL", gin.WrapH(
			playground.Handler("GraphQL", "/query"),
		))
	}

	// public routes
	r.GET("/", app.LandingPageHandler)
	r.GET("/projects/:id", app.ProjectHandler)
	r.Static("/static", app.staticFolder)
	r.POST("/login", app.loginHandler)

	// protected routes
	admin := app.AuthRequiredMiddleware()

	r.POST("/query", admin, gin.WrapH(srv))
	r.GET("/query", admin, gin.WrapH(srv))

	r.POST("/projects/:id/gallery/", admin, app.addProjectGallery)
	r.POST("/profile", admin, app.profileUploadHandler)
	r.POST("/resume", admin, app.resumeUploadHandler)
}
