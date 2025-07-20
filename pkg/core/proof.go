package core

type Proof struct {
	Total    int64    `bson:"total" json:"total"`
	Index    int64    `bson:"index" json:"index"`
	LeafHash string   `bson:"leafHash" json:"leafHash"` // hex encoded leaf hashes
	Aunts    []string `bson:"aunts" json:"aunts"`       // hex encoded aunts
}
