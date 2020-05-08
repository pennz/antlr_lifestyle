package model

import (
	"database/sql"

	"github.com/pennz/antlr_lifestyle/lifestyle"
)

type DataStore interface {
	AllThings() ([]*lifestyle.Thing, error)
	AllActions() ([]*lifestyle.Action, error)
	AllRelations() ([]*lifestyle.Relation, error)

	Thing(string) (*lifestyle.Thing, error)
	Action(string) (*lifestyle.Action, error)
	Relation(string) (*lifestyle.Relation, error)

	AddThing(*lifestyle.Thing) error
	AddAction(*lifestyle.Action) error
	AddRelation(*lifestyle.Relation) error
}

type DB struct {
	*sql.DB
}
