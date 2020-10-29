package lifestyle

import "time"

type Location string
type Why string
type How string
type Where string

type AnimalStatus struct { // we follow the ones in the game?
	Status
	Mood string
}

// Status a thing have a status, or status history, planned future status
type Status struct { // we follow the ones in the game?
	when  time.Time
	where Location
}
