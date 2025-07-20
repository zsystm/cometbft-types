package core

type Part struct {
	Index uint32 `bson:"index" json:"index"`
	Bytes string `bson:"bytes" json:"bytes"` // hex encoded bytes of the part
	Proof Proof  `bson:"proof" json:"proof"`
}
