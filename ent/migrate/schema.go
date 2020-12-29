// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true, SchemaType: map[string]string{"mysql": "varchar(32)", "postgres": "text", "sqlite3": "text"}},
		{Name: "email_address", Type: field.TypeString},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:        "accounts",
		Columns:     AccountsColumns,
		PrimaryKey:  []*schema.Column{AccountsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// CampaignsColumns holds the columns for the "campaigns" table.
	CampaignsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true, SchemaType: map[string]string{"mysql": "varchar(32)", "postgres": "text", "sqlite3": "text"}},
		{Name: "name", Type: field.TypeString},
		{Name: "organization_campaigns", Type: field.TypeUUID, Nullable: true},
	}
	// CampaignsTable holds the schema information for the "campaigns" table.
	CampaignsTable = &schema.Table{
		Name:       "campaigns",
		Columns:    CampaignsColumns,
		PrimaryKey: []*schema.Column{CampaignsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "campaigns_organizations_campaigns",
				Columns: []*schema.Column{CampaignsColumns[2]},

				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// OrganizationsColumns holds the columns for the "organizations" table.
	OrganizationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true, SchemaType: map[string]string{"mysql": "varchar(32)", "postgres": "text", "sqlite3": "text"}},
		{Name: "name", Type: field.TypeString},
	}
	// OrganizationsTable holds the schema information for the "organizations" table.
	OrganizationsTable = &schema.Table{
		Name:        "organizations",
		Columns:     OrganizationsColumns,
		PrimaryKey:  []*schema.Column{OrganizationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// OrganizationAccountsColumns holds the columns for the "organization_accounts" table.
	OrganizationAccountsColumns = []*schema.Column{
		{Name: "organization_id", Type: field.TypeUUID},
		{Name: "account_id", Type: field.TypeUUID},
	}
	// OrganizationAccountsTable holds the schema information for the "organization_accounts" table.
	OrganizationAccountsTable = &schema.Table{
		Name:       "organization_accounts",
		Columns:    OrganizationAccountsColumns,
		PrimaryKey: []*schema.Column{OrganizationAccountsColumns[0], OrganizationAccountsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "organization_accounts_organization_id",
				Columns: []*schema.Column{OrganizationAccountsColumns[0]},

				RefColumns: []*schema.Column{OrganizationsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:  "organization_accounts_account_id",
				Columns: []*schema.Column{OrganizationAccountsColumns[1]},

				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountsTable,
		CampaignsTable,
		OrganizationsTable,
		OrganizationAccountsTable,
	}
)

func init() {
	CampaignsTable.ForeignKeys[0].RefTable = OrganizationsTable
	OrganizationAccountsTable.ForeignKeys[0].RefTable = OrganizationsTable
	OrganizationAccountsTable.ForeignKeys[1].RefTable = AccountsTable
}
