package lifestyle

type thing struct {
	Status
	name string
}

type Actor struct {
	thing
}

type Receiver struct {
	thing
}

func (t *thing) StatusQuo() Status {
	return Status{}
}

// Act a action and change the status
func (t *thing) Act(Action) (Status, error) {
	return Status{}, nil
}

// changeStatus can only be done by action
func (t *thing) changeStatus(Status) (Status, error) {
	return Status{}, nil
}

func (t *thing) Relate(r Relation, other thing) {
}

func (t *thing) RelateMultiple(r Relation, others []thing) {
}
