package model

import (
	"database/sql"

	"gitlab.com/MrCue/antlr_lifestyle/lifestyle"
)

func AllActions(db *sql.DB) ([]*lifestyle.Action, error) {
	return nil, nil
}

func (d *DB) AllActions() ([]*lifestyle.Action, error) {
	rows, err := d.Query("SELECT * from action")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	actions := make([]*lifestyle.Action, 0)
	for rows.Next() {
		t := new(lifestyle.Action)
		err := rows.Scan(&(t.ActionType))
		if err != nil {
			return actions, err
		}
		actions = append(actions, t)
	}
	err = rows.Err()
	return actions, err
}

func (d *DB) Action(string) (*lifestyle.Action, error) {
	return nil, nil
}

func (d *DB) AddActionType(action_type lifestyle.ActionType) (int64, error) {
	stmt, err := d.Prepare("INSERT INTO action_type (name) VALUES($1)")
	if err != nil {
		return -1, err
	}
	_, err = stmt.Exec(action_type)
	if err != nil {
		return -1, err
	}
	return 1, nil
}

func (d *DB) AddAction(action *lifestyle.Action) error {
	debugType := "DEBUG"
	lastActionTypeId, err := d.AddActionType(lifestyle.ActionType(debugType))
	if err != nil {
		return err
	}

	stmt, err := d.Prepare("INSERT INTO action (name, action_type_id) VALUES($1, $2)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(action.ActionType, lastActionTypeId)
	if err != nil {
		return err
	}
	return nil
}
