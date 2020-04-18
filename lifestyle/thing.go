package lifestyle

type thingType struct { // meta info for Thing
	metaTypeName string
	domains      []thingDomain
}

func newthingDomainRegistry() *thingDomainRegistry {
	return &thingDomainRegistry{
		tangible:  "tangible",
		virtual:   "virtual",
		ephemeral: "ephemeral",
		eternal:   "eternal",
	}
}

type thingDomain string

type thingDomainRegistry struct {
	tangible  thingDomain
	virtual   thingDomain
	ephemeral thingDomain
	eternal   thingDomain
}

type Tag string

// Thing describes a thingship between things
type Thing struct {
	thingType
	name   string
	Status // now
	tags   []Tag
}

type Actor struct {
	Thing
}

type Receiver struct {
	Thing
}

func (t *Thing) StatusQuo() Status {
	return Status{}
}

// Act a action and change the status
func (t *Thing) Act(Action) (Status, error) {
	return Status{}, nil
}

// changeStatus can only be done by action
func (t *Thing) changeStatus(Status) (Status, error) {
	return Status{}, nil
}

func (t *Thing) Relate(r Relation, other Thing) {
}

func (t *Thing) RelateMultiple(r Relation, others []Thing) {

}
