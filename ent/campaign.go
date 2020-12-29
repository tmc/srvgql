// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/tmc/pulid"
	"github.com/tmc/srvgql/ent/campaign"
)

// Campaign is the model entity for the Campaign schema.
type Campaign struct {
	config `json:"-"`
	// ID of the ent.
	ID pulid.ID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name                   string `json:"name,omitempty"`
	organization_campaigns *pulid.ID
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Campaign) scanValues() []interface{} {
	return []interface{}{
		&pulid.ID{},       // id
		&sql.NullString{}, // name
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Campaign) fkValues() []interface{} {
	return []interface{}{
		&pulid.ID{}, // organization_campaigns
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Campaign fields.
func (c *Campaign) assignValues(values ...interface{}) error {
	if m, n := len(values), len(campaign.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*pulid.ID); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value != nil {
		c.ID = *value
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		c.Name = value.String
	}
	values = values[1:]
	if len(values) == len(campaign.ForeignKeys) {
		if value, ok := values[0].(*pulid.ID); !ok {
			return fmt.Errorf("unexpected type %T for field organization_campaigns", values[0])
		} else if value != nil {
			c.organization_campaigns = value
		}
	}
	return nil
}

// Update returns a builder for updating this Campaign.
// Note that, you need to call Campaign.Unwrap() before calling this method, if this Campaign
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Campaign) Update() *CampaignUpdateOne {
	return (&CampaignClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Campaign) Unwrap() *Campaign {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Campaign is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Campaign) String() string {
	var builder strings.Builder
	builder.WriteString("Campaign(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", name=")
	builder.WriteString(c.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Campaigns is a parsable slice of Campaign.
type Campaigns []*Campaign

func (c Campaigns) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}