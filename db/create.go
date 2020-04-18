package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	//_ import driver
	_ "github.com/lib/pq"
)

const (
	THING                     = "thing"
	THING_TYPE                = "thing_type"
	THING_DOMAIN              = "thing_domain"
	TAG                       = "tag"
	STATUS                    = "status"
	ACTION                    = "action"
	ACTION_TYPE               = "action_type"
	RELATION                  = "relation"
	RELATION_TYPE             = "relation_type"
	RELATION_META_TYPE        = "relation_meta_type"
	RELATION_META_AMOUNT_TYPE = "relation_meta_amount_type"
)

func Create() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	// we need to write our own sql create script
	// we have these data : relation, action, thing, status
	// relation have relation_meta, relation, it is related to thing
	// action have actionRegistry, it should be add and fill in the way
	// ting has status, have actions, relations
	// just do this first
	sqlCreateStr := `
CREATE TABLE thing (
  id SERIAL,
  name varchar(255),
  status_id,
  thing_type_id
);
CREATE TABLE thing_type (
  id SERIAL,
  name varchar(255),
);
CREATE TABLE thing_domain (
  id SERIAL,
  name varchar(255)
);
CREATE TABLE status (
  id SERIAL,
);
CREATE TABLE tag (
  id SERIAL,
  name varchar(255)
);
CREATE TABLE action (
  id SERIAL,
  name varchar(255)
);
CREATE TABLE action_type (
  id SERIAL,
  name varchar(255)
); from to why how
CREATE TABLE relation (
  id SERIAL,
  name varchar(255)
);
CREATE TABLE relation_type (
  id SERIAL,
  name varchar(255)
);
CREATE TABLE relation_meta_type (
  id SERIAL,
  reciprocal boolean,
  meta_name varchar(255)
);
CREATE TABLE relation_meta_amount_type (
  id SERIAL,
  name varchar(255)
);
`

	tablesWithIndex := []string{THING, THING_TYPE, THING_DOMAIN, TAG}
	tablesFKs := map[string]([]string){
		THING:              []string{THING_TYPE},
		STATUS:             []string{THING},
		RELATION:           []string{RELATION_TYPE},
		RELATION_TYPE:      []string{RELATION_META_TYPE},
		RELATION_META_TYPE: []string{RELATION_META_AMOUNT_TYPE}, // make it easy, just 1 to 1...
		ACTION:             []string{ACTION_TYPE},
	}

	tablesM2M := map[string]([]string){
		THING:      []string{TAG, ACTION, RELATION}, // for action, just 1 to 1 for now
		THING_TYPE: []string{THING_DOMAIN},
		RELATION:   []string{THING},
	}

	for _, t := range tablesWithIndex {
		sqlCreateStr += fmt.Sprintf("ALTER TABLE %s ADD PRIMARY KEY (id);\n", t)
	}
	for tMany, tOnes := range tablesFKs {
		for _, tOne := range tOnes {
			sqlCreateStr += fmt.Sprintf("ALTER TABLE %s ADD %s_id int NOT NULL;\n", tMany, tOne)
			sqlCreateStr += fmt.Sprintf("ALTER TABLE %s ADD FOREIGN KEY (%s_id) REFERENCES %s (id);\n", tMany, tOne, tOne)
		}
	}
	for m1, m2s := range tablesM2M {
		for _, m2 := range m2s {
			sqlCreateStr += fmt.Sprintf(`
CREATE TABLE %s_%s (
  id SERIAL,
  %s_id int NOT NULL,
  %s_id int NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (%s_id) REFERENCES %s(id),
  FOREIGN KEY (%s_id) REFERENCES %s(id)
);\n`, m1, m2, m1, m2, m1, m1, m2, m2)
		}
	}

	log.Println(sqlCreateStr)
	log.Println(db)
	defer db.Close()
}
