package main

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/tmc/srvgql/ent"
)

func TestORM(t *testing.T) {
	ctx := context.Background()
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Create account
	u, err := CreateAccount(ctx, client)
	if err != nil {
		log.Println(err)
	}
	o, err := CreateOrganization(ctx, client, u)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(u)
	fmt.Println(o)

	a, err := QueryAccount(ctx, client)
	if err != nil {
		log.Println("err querying accounts:", err)
	}
	log.Println("account", a)
	log.Println("orgs", a.Edges.Organization)
}
