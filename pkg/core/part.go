package core

type Part struct {
	Index uint32 `bson:"index"`
	Bytes string `bson:"bytes"` // hex encoded bytes of the part
	Proof Proof  `bson:"proof"`
}
