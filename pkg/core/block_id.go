package core

type BlockID struct {
	Hash          string        `bson:"hash" json:"hash"`
	PartSetHeader PartSetHeader `bson:"partSetHeader" json:"partSetHeader"`
}

type PartSetHeader struct {
	Total uint64 `bson:"total" json:"total"`
	Hash  string `bson:"hash" json:"hash"`
}
