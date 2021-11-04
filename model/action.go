package model

import (
	"database/sql"

	"gitlab.com/MrCue/antlr_lifestyle/lifestyle"
)

func AllActions(db *sql.DB) ([]*lifestyle.Action, error) {
	return nil, nil
}

func (d *DB) AllActions() ([]*lifestyle.Action, error) {
	return nil, nil
}
func (d *DB) Action(string) (*lifestyle.Action, error) {
	return nil, nil
}
func (d *DB) AddAction(*lifestyle.Action) error {
	return nil
}
