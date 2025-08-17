//go:generate go run github.com/99designs/gqlgen generate
package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/openai/openai-go/v2"

	"github.com/joho/godotenv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jiten-mobile/service/core"
	"github.com/jiten-mobile/service/graph"
	"github.com/vektah/gqlparser/v2/ast"
)

var firestoreClient *firestore.Client
var openaiClient *openai.Client

func graphqlHandler() gin.HandlerFunc {
	h := handler.New(graph.NewExecutableSchema(
		graph.Config{Resolvers: graph.NewDictionaryResolver(firestoreClient, openaiClient)}),
	)

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
	ctx := context.Background()
	firestoreClient, _ = core.InitFirestoreClient(ctx)
	openaiClient = core.InitOpenaiClient()
}

func main() {
	r := gin.Default()
	defer firestoreClient.Close()
	r.POST("/v1/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run()
}
