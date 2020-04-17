package lifestyle

var Actions = newactionRegistry()

type ActionType string

//https://www.linguasorb.com/english/verbs/most-common-verbs/
func newactionRegistry() *actionRegistry {
	return &actionRegistry{
		Be:         "be",
		Have:       "have",
		Do:         "do",
		Say:        "say",
		Go:         "go",
		Can:        "can",
		Get:        "get",
		Would:      "would",
		Make:       "make",
		Know:       "know",
		Will:       "will",
		Think:      "think",
		Take:       "take",
		See:        "see",
		Come:       "come",
		Could:      "could",
		Want:       "want",
		Look:       "look",
		Use:        "use",
		Find:       "find",
		Give:       "give",
		Tell:       "tell",
		Work:       "work",
		May:        "may",
		Should:     "should",
		Call:       "call",
		Try:        "try",
		Ask:        "ask",
		Need:       "need",
		Feel:       "feel",
		Become:     "become",
		Leave:      "leave",
		Put:        "put",
		Mean:       "mean",
		Keep:       "keep",
		Let:        "let",
		Begin:      "begin",
		Seem:       "seem",
		Help:       "help",
		Talk:       "talk",
		Turn:       "turn",
		Start:      "start",
		Might:      "might",
		Show:       "show",
		Hear:       "hear",
		Play:       "play",
		Run:        "run",
		Move:       "move",
		Like:       "like",
		Live:       "live",
		Believe:    "believe",
		Hold:       "hold",
		Bring:      "bring",
		Happen:     "happen",
		Must:       "must",
		Write:      "write",
		Provide:    "provide",
		Sit:        "sit",
		Stand:      "stand",
		Lose:       "lose",
		Pay:        "pay",
		Meet:       "meet",
		Include:    "include",
		Continue:   "continue",
		Set:        "set",
		Learn:      "learn",
		Change:     "change",
		Lead:       "lead",
		Understand: "understand",
		Watch:      "watch",
		Follow:     "follow",
		Stop:       "stop",
		Create:     "create",
		Speak:      "speak",
		Read:       "read",
		Allow:      "allow",
		Add:        "add",
		Spend:      "spend",
		Grow:       "grow",
		Open:       "open",
		Walk:       "walk",
		Win:        "win",
		Offer:      "offer",
		Remember:   "remember",
		Love:       "love",
		Consider:   "consider",
		Appear:     "appear",
		Buy:        "buy",
		Wait:       "wait",
		Serve:      "serve",
		Die:        "die",
		Send:       "send",
		Expect:     "expect",
		Build:      "build",
		Stay:       "stay",
		Fall:       "fall",
		Cut:        "cut",
		Reach:      "reach",
		Kill:       "kill",
		Remain:     "remain",
	}
}

type actionRegistry struct {
	Be         ActionType
	Have       ActionType
	Do         ActionType
	Say        ActionType
	Go         ActionType
	Can        ActionType
	Get        ActionType
	Would      ActionType
	Make       ActionType
	Know       ActionType
	Will       ActionType
	Think      ActionType
	Take       ActionType
	See        ActionType
	Come       ActionType
	Could      ActionType
	Want       ActionType
	Look       ActionType
	Use        ActionType
	Find       ActionType
	Give       ActionType
	Tell       ActionType
	Work       ActionType
	May        ActionType
	Should     ActionType
	Call       ActionType
	Try        ActionType
	Ask        ActionType
	Need       ActionType
	Feel       ActionType
	Become     ActionType
	Leave      ActionType
	Put        ActionType
	Mean       ActionType
	Keep       ActionType
	Let        ActionType
	Begin      ActionType
	Seem       ActionType
	Help       ActionType
	Talk       ActionType
	Turn       ActionType
	Start      ActionType
	Might      ActionType
	Show       ActionType
	Hear       ActionType
	Play       ActionType
	Run        ActionType
	Move       ActionType
	Like       ActionType
	Live       ActionType
	Believe    ActionType
	Hold       ActionType
	Bring      ActionType
	Happen     ActionType
	Must       ActionType
	Write      ActionType
	Provide    ActionType
	Sit        ActionType
	Stand      ActionType
	Lose       ActionType
	Pay        ActionType
	Meet       ActionType
	Include    ActionType
	Continue   ActionType
	Set        ActionType
	Learn      ActionType
	Change     ActionType
	Lead       ActionType
	Understand ActionType
	Watch      ActionType
	Follow     ActionType
	Stop       ActionType
	Create     ActionType
	Speak      ActionType
	Read       ActionType
	Allow      ActionType
	Add        ActionType
	Spend      ActionType
	Grow       ActionType
	Open       ActionType
	Walk       ActionType
	Win        ActionType
	Offer      ActionType
	Remember   ActionType
	Love       ActionType
	Consider   ActionType
	Appear     ActionType
	Buy        ActionType
	Wait       ActionType
	Serve      ActionType
	Die        ActionType
	Send       ActionType
	Expect     ActionType
	Build      ActionType
	Stay       ActionType
	Fall       ActionType
	Cut        ActionType
	Reach      ActionType
	Kill       ActionType
	Remain     ActionType
}

type Location string
type Time string
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

type Why string
type How string

type Action struct {
	Actor      // where who
	to         Receiver
	ActionType // when
	Why
	How
}

type AnimalStatus struct { // we follow the ones in the game?
	Status
	Mood string
}

// Status a thing have a status, or status history, planned future status
type Status struct { // we follow the ones in the game?
	when  Time
	where Location
}

// NewAction  who where when what why how
func NewAction() *Action {
	return &Action{Actor{}, Receiver{}, Actions.Be, "", ""}
}
