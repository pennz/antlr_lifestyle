package lifestyle

import "fmt"

var Actions = newactionRegistry()

type ActionType string

func newactionRegistry() *actionRegistry {
	return &actionRegistry{
		Red:   "red",
		Green: "green",
		Blue:  "blue",
	}
}

type actionRegistry struct {
	Red   ActionType
	Green ActionType
	Blue  ActionType
}

type Location string
type Time string
type thing struct {
	where Location
	name  string
	when  Time
}

type Actor struct {
	thing
}

type Receiver struct {
	thing
}

type Why string
type How string

type Action struct {
	Actor      // where who
	to         Receiver
	ActionType // when
	Why
	How
}

//  who where when what why how

func main() {
	fmt.Println("vim-go")
	fmt.Println(Actions.Red)
}
