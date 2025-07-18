package core

type Proof struct {
	Total    int64    `bson:"total"`
	Index    int64    `bson:"index"`
	LeafHash string   `bson:"leafHash"` // hex encoded leaf hashes
	Aunts    []string `bson:"aunts"`    // hex encoded aunts
}
