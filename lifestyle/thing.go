package lifestyle

type thingType struct { // meta info for Thing
	metaTypeName string
	thingDomain
}

func newthingDomainRegistry() *thingDomainRegistry {
	return &thingDomainRegistry{
		One2One:   "One2One",
		One2Many:  "One2Many",
		Many2One:  "Many2One",
		Many2Many: "Many2Many",
	}
}

type thingDomain string

type thingDomainRegistry struct {
	One2One   thingDomain
	One2Many  thingDomain
	Many2One  thingDomain
	Many2Many thingDomain
}

type Tag string

// Thing describes a thingship between things
type Thing struct {
	thingType
	name string
	Status
	tags []Tag
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
