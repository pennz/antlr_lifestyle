package lifestyle

type relationType struct { // meta info for relation
	metaTypeName string
	reciprocal   bool          // (1<->1) (*-*)
	amount       relamountType // 1-1, 1-*, *-1, *-*, 1-0?
}

func newrelamountRegistry() *relamountRegistry {
	return &relamountRegistry{
		One2One:   "One2One",
		One2Many:  "One2Many",
		Many2One:  "Many2One",
		Many2Many: "Many2Many",
	}
}

type relamountType string

type relamountRegistry struct {
	One2One   relamountType
	One2Many  relamountType
	Many2One  relamountType
	Many2Many relamountType
}

// Relation describes a relationship between things
type Relation struct {
	relationType
	name string
}
