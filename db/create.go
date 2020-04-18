package db

import (
	"context"
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

var (
	ctx context.Context
	db  *sql.DB
)

func Create() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	//if err := db.PingContext(ctx); err != nil {
	//	log.Fatal(err)
	//}

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
  PRIMARY KEY (id)
);
CREATE TABLE thing_type (
  id SERIAL,
  name varchar(255),
  PRIMARY KEY (id)
);
CREATE TABLE thing_domain (
  id SERIAL,
  name varchar(255),
  PRIMARY KEY (id)
);
CREATE TABLE status (
  id SERIAL,
  PRIMARY KEY (id)
);
CREATE TABLE tag (
  id SERIAL,
  name varchar(255),
  PRIMARY KEY (id)
);
CREATE TABLE action (
  id SERIAL,
  name varchar(255),
  PRIMARY KEY (id)
);
CREATE TABLE action_type (
  id SERIAL,
  name varchar(255),
  PRIMARY KEY (id)
);
CREATE TABLE relation (
  id SERIAL,
  name varchar(255),
  from_id int NOT NULL,
  to_id int NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (from_id) REFERENCES thing(id),
  FOREIGN KEY (to_id)   REFERENCES thing(id)
);
CREATE TABLE relation_type (
  id SERIAL,
  name varchar(255),
  PRIMARY KEY (id)
);
CREATE TABLE relation_meta_type (
  id SERIAL,
  reciprocal boolean,
  meta_name varchar(255),
  PRIMARY KEY (id)
);
CREATE TABLE relation_meta_amount_type (
  id SERIAL,
  name varchar(255),
  PRIMARY KEY (id)
);
`

	tablesFKs := map[string]([]string){
		THING:              []string{THING_TYPE},
		STATUS:             []string{THING},
		RELATION:           []string{RELATION_TYPE},
		RELATION_TYPE:      []string{RELATION_META_TYPE},
		RELATION_META_TYPE: []string{RELATION_META_AMOUNT_TYPE}, // make it easy, just 1 to 1...
		ACTION:             []string{ACTION_TYPE},
	}

	tablesM2M := map[string]([]string){
		THING:      []string{TAG, ACTION}, // for action, just 1 to 1 for now
		THING_TYPE: []string{THING_DOMAIN},
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
);`, m1, m2, m1, m2, m1, m1, m2, m2)
		}
	}
	// for relation ,it just exist in relation db
	log.Println(sqlCreateStr)
	result, err := db.Exec(sqlCreateStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)

	log.Println(sqlCreateStr)
	log.Println(db)
	defer db.Close()
}
