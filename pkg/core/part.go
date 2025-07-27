package core

import (
	"crypto/sha256"
	"strconv"
)

type Part struct {
	Index uint32 `bson:"index" json:"index"`
	Bytes string `bson:"bytes" json:"bytes"` // hex encoded bytes of the part
	Proof Proof  `bson:"proof" json:"proof"`
}

func (p Part) Hash() []byte {
	h := sha256.New()
	h.Write([]byte(strconv.FormatUint(uint64(p.Index), 10)))
	h.Write([]byte(p.Bytes))
	h.Write(p.Proof.Hash())
	return h.Sum(nil)
}
