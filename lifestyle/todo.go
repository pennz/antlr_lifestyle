package lifestyle

type CalenderItem interface {
	AddReminder(When)
	Countdown() Time
}

// ToDo actions change the world.
type ToDo interface {
	CalenderItem
	Manage
	Tagee
	ToDoRaw
}

type Manage interface {
	Logging
	Review() string
	Exalate()
	Suspend()
	Continue()
	AddMilestone()
	CheckMilestone()
}

type Logging interface {
	AddResourceUsage()
	Serialize() // TODO find existing type
	GetProcess()
}

// Tag for tag data classification, for class, we can use tag too
type Tag string
type Tagee interface {
	AddTag(Tag)
	RemoveTag(Tag)
}

// DataOrganizer show your data in an organized way.
type DataOrganizer interface {
	Filter(Tag) []Tagee
	Export()
	Import()
	Sort()
}

type ToDoRaw interface {
	NewToDo() *ToDo
	Do() Result
	AddSteps()
	SetDeadLine()
	AddRelated(Thing)
}

func NewToDo() *Action {
	a := Action{ActionType: Actions.Will}
	return &a
}
