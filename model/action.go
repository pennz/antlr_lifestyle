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
func (d *DB) AddAction(action *lifestyle.Action) error {
	stmt, err := d.Prepare("INSERT INTO action(name) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(action.ActionType)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
}
