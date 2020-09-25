package lifestyle

import (
	"log"
	"time"
)

// ToBe existence
type ToBe struct {
	*log.Logger
}

func (t *ToBe) AddReminder(when time.Time) {
	t.Logger.Print(when)
}

func (t *ToBe) Countdown() time.Time {
	panic("not implemented") // TODO: Implement
}

func (t *ToBe) AddResourceUsage() {
	panic("not implemented") // TODO: Implement
}

func (t *ToBe) Serialize() {
	panic("not implemented") // TODO: Implement
}

func (t *ToBe) GetProcess() {
	panic("not implemented") // TODO: Implement
}

func (t *ToBe) Review() string {
	panic("not implemented") // TODO: Implement
}

func (t *ToBe) Exalate() {
	panic("not implemented") // TODO: Implement
}

func (t *ToBe) Suspend() {
	panic("not implemented") // TODO: Implement
}

func (t *ToBe) Continue() {
	panic("not implemented") // TODO: Implement
}

func (t *ToBe) AddMilestone() {
	panic("not implemented") // TODO: Implement
}

func (t *ToBe) CheckMilestone() {
	panic("not implemented") // TODO: Implement
}

func (t *ToBe) AddTag(_ Tag) {
	panic("not implemented") // TODO: Implement
}

func (t *ToBe) RemoveTag(_ Tag) {
	panic("not implemented") // TODO: Implement
}

func (t *ToBe) NewToDo() *ToDo {
	panic("not implemented") // TODO: Implement
}

func (t *ToBe) Do() Result {
	panic("not implemented") // TODO: Implement
}

func (t *ToBe) AddSteps() {
	panic("not implemented") // TODO: Implement
}

func (t *ToBe) SetDeadLine() {
	panic("not implemented") // TODO: Implement
}

func (t *ToBe) AddRelated(_ Thing) {
	panic("not implemented") // TODO: Implement
}
