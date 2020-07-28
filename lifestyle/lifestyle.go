package lifestyle

type Location string
type Time string
type Why string
type How string
type Where string
type When string

type AnimalStatus struct { // we follow the ones in the game?
	Status
	Mood string
}

// Status a thing have a status, or status history, planned future status
type Status struct { // we follow the ones in the game?
	when  Time
	where Location
}
