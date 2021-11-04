package model

import (
	"database/sql"
	"log"

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
func (d *DB) AddAction(action *lifestyle.Action) error {
	stmt, err := d.Prepare("INSERT INTO action(name) VALUES(?)")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(action.ActionType)
	if err != nil {
		return err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
	return nil
}
