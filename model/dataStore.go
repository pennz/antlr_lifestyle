package model

import (
	"database/sql"

	"github.com/pennz/antlr_lifestyle/lifestyle"
)

type DataStore interface {
	AllThings() ([]*lifestyle.Thing, error)
	AllActions() ([]*lifestyle.Action, error)
	AllRelations() ([]*lifestyle.Relation, error)
}

type DB struct {
	*sql.DB
}
