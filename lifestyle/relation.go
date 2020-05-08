package lifestyle

type relationMetaType struct {
	metaName   string
	reciprocal bool                   // (1<->1) (*-*)
	amount     relationMetaAmountType // 1-1, 1-*, *-1, *-*, 1-0?
}
type relationType struct { // meta info for relation
	name string
	relationMetaType
}

func newrelationMetaAmountRegistry() *relationMetaAmountRegistry {
	return &relationMetaAmountRegistry{
		One2One:   "One2One",
		One2Many:  "One2Many",
		Many2One:  "Many2One",
		Many2Many: "Many2Many",
	}
}

type relationMetaAmountType string

type relationMetaAmountRegistry struct {
	One2One   relationMetaAmountType
	One2Many  relationMetaAmountType
	Many2One  relationMetaAmountType
	Many2Many relationMetaAmountType
}

// Relation describes a relationship between things
// it is many-to-many
type Relation struct {
	relationType
	Name string
	from []Thing // just for datastructure, the relation type will enforce amount
	to   []Thing
}
