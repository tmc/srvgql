package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/srvgql/ent"
)

func CreateAccount(ctx context.Context, client *ent.Client) (*ent.Account, error) {
	u, err := client.Account.
		Create().
		SetEmailAddress("test@testco.us").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating account: %v", err)
	}
	log.Println("account was created: ", u)
	return u, nil
}

func CreateOrganization(ctx context.Context, client *ent.Client, account *ent.Account) (*ent.Organization, error) {
	o, err := client.Organization.
		Create().
		SetName("testco, Inc.").
		AddAccounts(account).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating account: %v", err)
	}
	log.Println("org was created: ", o)
	return o, nil
}
