package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jiten-mobile/service/db"
	"github.com/jiten-mobile/service/graph"
	"github.com/vektah/gqlparser/v2/ast"
)

func graphqlHandler() gin.HandlerFunc {
	h := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	h.AddTransport(transport.Options{})
	h.AddTransport(transport.GET{})
	h.AddTransport(transport.POST{})

	h.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	h.Use(extension.Introspection{})
	h.Use(extension.AutomaticPersistedQuery{Cache: lru.New[string](100)})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/v1/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading env file %v", err)
	}
	db.InitClient()
}

func main() {
	r := gin.Default()
	r.POST("/v1/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run()
}
