package core

type BitArray struct {
	Bits  int64    `bson:"bits" json:"bits"`
	Elems []uint64 `bson:"elems" json:"elems"`
}
