package model

import (
	"database/sql"

	"github.com/pennz/antlr_lifestyle/lifestyle"
)

// some version of data access, needs improve
func AllThings(db *sql.DB) ([]*lifestyle.Thing, error) {
	rows, err := db.Query("SELECT * from thing")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	things := make([]*lifestyle.Thing, 0)
	for rows.Next() {
		t := new(lifestyle.Thing)
		err := rows.Scan(&lifestyle.Thing.Name)
		if err != nil {
			return things, err
		}
		things = append(things, t)
	}
	err = rows.Err()
	return things, err

}
