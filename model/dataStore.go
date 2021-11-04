package model

import (
	"database/sql"

	"gitlab.com/MrCue/antlr_lifestyle/lifestyle"
)

type DB struct {
	*sql.DB // https://stackoverflow.com/questions/44406077/golang-invalid-receiver-type-in-method-func
}

// DataStore for our website
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
