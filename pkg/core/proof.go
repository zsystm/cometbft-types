package core

import (
	"crypto/sha256"
	"strconv"
)

type Proof struct {
	Total    int64    `bson:"total" json:"total"`
	Index    int64    `bson:"index" json:"index"`
	LeafHash string   `bson:"leafHash" json:"leafHash"` // hex encoded leaf hashes
	Aunts    []string `bson:"aunts" json:"aunts"`       // hex encoded aunts
}

func (p *Proof) Hash() []byte {
	h := sha256.New()
	h.Write([]byte(strconv.FormatInt(p.Index, 10)))
	h.Write([]byte(strconv.FormatInt(p.Total, 10)))
	h.Write([]byte(p.LeafHash))
	for _, aunt := range p.Aunts {
		h.Write([]byte(aunt))
	}
	return h.Sum(nil)
}
