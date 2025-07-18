package core

type BitArray struct {
	Bits  int64    `bson:"bits"`
	Elems []uint64 `bson:"elems"`
}
