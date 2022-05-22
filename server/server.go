package main

import (
	"github.com/DedAzaMarks/hw07-graphql/server/repository"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/DedAzaMarks/hw07-graphql/server/graph"
	"github.com/DedAzaMarks/hw07-graphql/server/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	players := []repository.Player{
		{
			Username:  "Player1",
			Role:      repository.Mafia,
			IsAlive:   true,
			DaysAlive: 1,
		},
		{
			Username:  "Player2",
			Role:      repository.Civilian,
			IsAlive:   true,
			DaysAlive: 1,
		},
		{
			Username:  "Player3",
			Role:      repository.Civilian,
			IsAlive:   false,
			DaysAlive: 0,
		},
	}
	repo := repository.Repository{
		Players: players,
		Games: []repository.Game{
			{
				Id:         uuid.New(),
				Scoreboard: players,
				Comments:   []string{"no comments"},
			},
		},
	}

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: graph.NewResolver(repo)}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
