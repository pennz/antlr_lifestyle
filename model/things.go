package model

import (
	"database/sql"

	"gitlab.com/MrCue/antlr_lifestyle/lifestyle"
)

// follow this tutorial https://www.alexedwards.net/blog/organising-database-access
// some version of data access, needs improve
// use db as a global variable in the model package, then
// we can use that to all the data retrieval
func AllThings(db *sql.DB) ([]*lifestyle.Thing, error) {
	rows, err := db.Query("SELECT * from thing")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	things := make([]*lifestyle.Thing, 0)
	for rows.Next() {
		t := new(lifestyle.Thing)
		err := rows.Scan(&(t.Name))
		if err != nil {
			return things, err
		}
		things = append(things, t)
	}
	err = rows.Err()
	return things, err

}

// 2. dependency injection
// we use a Env struct store all our database related thing
// then function will be related to that struct.
