package main

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/srvgql/ent"
	"github.com/tmc/srvgql/ent/account"
)

func QueryAccount(ctx context.Context, client *ent.Client) (*ent.Account, error) {
	u, err := client.Account.
		Query().
		WithOrganization().
		Where(account.EmailAddressEQ("test@testco.us")).
		// `Only` fails if no account found,
		// or more than 1 account returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying account: %v", err)
	}
	log.Println("account returned: ", u)
	return u, nil
}
