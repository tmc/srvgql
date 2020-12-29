package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tmc/srvgql/ent"
	"github.com/tmc/srvgql/graph"
)

var (
	flagDatabaseType = flag.String("db-type", "sqlite3", "type of database (one of sqlite3, mysql, postgres)")
	flagDatabaseURL  = flag.String("db-url", "file:ent.sqlite?_fk=1", "DATABASE_URL")
	flagVerbose      = flag.Bool("v", false, "verbose mode")
)

func main() {
	flag.Parse()
	ctx := context.Background()

	client, err := ent.Open(*flagDatabaseType, *flagDatabaseURL)
	if *flagVerbose {
		client = client.Debug()
	}
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	if err := seedBasicTestData(ctx, client); err != nil {
		log.Fatal("issue seeding test data:", err)
	}

	ctx = ent.NewContext(ctx, client)
	if err := serveGQL(ctx); err != nil {
		log.Fatal("issue serving gql service:", err)
	}
}

func seedBasicTestData(ctx context.Context, client *ent.Client) error {
	// Create account
	u, err := CreateAccount(ctx, client)
	if err != nil {
		return fmt.Errorf("issue creating account organization: %w", err)
	}
	o, err := CreateOrganization(ctx, client, u)
	if err != nil {
		return fmt.Errorf("issue creating organization: %w", err)
	}
	fmt.Println(u)
	fmt.Println(o)

	a, err := QueryAccount(ctx, client)
	if err != nil {
		return fmt.Errorf("issue querying: %w", err)
	}
	log.Println("account", a)
	log.Println("orgs", a.Edges.Organization)
	return nil
}

func serveGQL(ctx context.Context) error {
	defaultPort := "8080"
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Client: ent.FromContext(ctx),
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	return http.ListenAndServe(":"+port, nil)
}
