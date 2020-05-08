package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	//_ import driver
	_ "github.com/lib/pq"
	"github.com/pennz/antlr_lifestyle/lifestyle"
)

type Env struct {
	DS DataStore
}

var dataEnv Env

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

type fakeDB struct{}

func (fdb *fakeDB) AllThings() ([]*lifestyle.Thing, error) {
	things := make([]*lifestyle.Thing, 0)
	things = append(things, &lifestyle.Thing{Name: "No"})
	things = append(things, &lifestyle.Thing{Name: "Yes"})
	return things, nil
}
func (fdb *fakeDB) AllActions() ([]*lifestyle.Action, error) {
	actions := make([]*lifestyle.Action, 0)
	actions = append(actions, &lifestyle.Action{ActionType: lifestyle.Actions.Run})
	actions = append(actions, &lifestyle.Action{ActionType: lifestyle.Actions.Run})
	return actions, nil
}

func (fdb *fakeDB) Action(string) (*lifestyle.Action, error) {
	return nil, nil
}
func (fdb *fakeDB) Relation(string) (*lifestyle.Relation, error) { return nil, nil }
func (fdb *fakeDB) Thing(string) (*lifestyle.Thing, error)       { return nil, nil }

func (fdb *fakeDB) AllRelations() ([]*lifestyle.Relation, error) {
	relations := make([]*lifestyle.Relation, 0)
	relations = append(relations, &lifestyle.Relation{Name: "No"})
	relations = append(relations, &lifestyle.Relation{Name: "Yes"})
	return relations, nil
}

func (fdb *fakeDB) AddThing(*lifestyle.Thing) error       { return nil }
func (fdb *fakeDB) AddAction(*lifestyle.Action) error     { return nil }
func (fdb *fakeDB) AddRelation(*lifestyle.Relation) error { return nil }

// GetFakeEnv get it for test
func GetFakeEnv() (Env, error) {
	return Env{&fakeDB{}}, nil
}

// GetDB is like factory, create new one if not there
func GetDB() (Env, error) {
	var err error = nil
	if dataEnv.DS == nil {
		db, err := getDB()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(db)
		dataEnv = Env{&fakeDB{}}
	}

	return dataEnv, err
}

func getDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}

func create() *sql.DB {
	//if err := db.PingContext(ctx); err != nil {
	//	log.Fatal(err)
	//}
	db, err := getDB()
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

	return db
}
