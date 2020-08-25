package lifestyle

type ToBe struct {
}

// ToDo actions change the world.
type ToDo interface {
	_ToDo
	logging
	calenderItem
	manage
	tagee
}

type manage interface {
	logging
	Review() string
	Exalate()
	Suspend()
	Continue()
	AddMilestone()
	CheckMilestone()
}

type calenderItem interface {
	AddReminder(When)
	Countdown() Time
}

type logging interface {
	AddResourceUsage()
	Serialize() // TODO find existing type
	GetProcess()
}

// Tag for tag data classification, for class, we can use tag too
type Tag string
type tagee interface {
	AddTag(Tag)
	RemoveTag(Tag)
}

// DataOrganizer show your data in an organized way.
type DataOrganizer interface {
	Filter(Tag) []tagee
	Export()
	Import()
	Sort()
}

type _ToDo interface {
	NewToDo() *ToDo
	Do() Result
	AddSteps()
	SetDeadLine()
	AddRelated(Thing)
}

func newToDo() *Action {
	a := Action{ActionType: Actions.Will}
	return &a
}
